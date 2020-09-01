package gooMQ

import (
	"fmt"
	"github.com/Shopify/sarama"
	gooLog "googo.io/goo/log"
	"log"
	"time"
)

type KafkaProducer struct {
	Kafka
	producer sarama.AsyncProducer
}

func (this *KafkaProducer) getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = 5 * time.Second
	config.Version = sarama.V0_10_2_0
	return config
}

func (this *KafkaProducer) Init() {
	producer, err := sarama.NewAsyncProducer(this.Addrs, this.getConfig())
	if err != nil {
		gooLog.Error("[kafka-producer-error]", err.Error())
		log.Panic(err.Error())
	}

	this.producer = producer

	go func() {
		for {
			select {
			case suc := <-this.producer.Successes():
				gooLog.Debug("[kafka-async-producer-succ]", fmt.Sprintf("partitions=%d topic=%s offset=%d value=%s",
					suc.Partition, suc.Topic, suc.Offset, suc.Value))
			case err := <-this.producer.Errors():
				gooLog.Error("[kafka-async-producer-error]", err.Error())
			case <-this.Context.Done():
				return
			}
		}
	}()
}

func (this *KafkaProducer) SendMessage(topic string, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
		Key:   sarama.StringEncoder(fmt.Sprintf("%d", time.Now().UnixNano())),
	}

	this.producer.Input() <- msg

	return nil
}
