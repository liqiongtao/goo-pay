package gooMQ

import (
	"fmt"
	"github.com/Shopify/sarama"
	gooLog "googo.io/goo/log"
	"sync"
	"time"
)

type KafkaProducer struct {
	*Kafka
	producer sarama.AsyncProducer
	rwmu     sync.RWMutex
}

func (this *KafkaProducer) config() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = 5 * time.Second
	config.Version = sarama.V0_10_2_0
	return config
}

func (this *KafkaProducer) setProducer() sarama.AsyncProducer {
	this.rwmu.Lock()
	defer this.rwmu.Unlock()

	if this.producer != nil {
		return this.producer
	}

	producer, err := sarama.NewAsyncProducer(this.Addrs, this.config())
	if err != nil {
		gooLog.Error("[kafka-producer-error]", err.Error())
		panic(err.Error())
	}

	go func() {
		for {
			select {
			case suc := <-producer.Successes():
				gooLog.Debug("[kafka-async-producer-succ]",
					fmt.Sprintf("partitions=%d topic=%s offset=%d value=%s",
						suc.Partition, suc.Topic, suc.Offset, suc.Value))
				
			case err := <-producer.Errors():
				gooLog.Error("[kafka-async-producer-error]", err.Error())

			case <-this.Context.Done():
				return
			}
		}
	}()

	this.producer = producer
	return this.producer
}

func (this *KafkaProducer) getProducer() sarama.AsyncProducer {
	this.rwmu.RLock()
	defer this.rwmu.RUnlock()

	return this.producer
}

func (this *KafkaProducer) SendMessage(topic string, value []byte) error {
	if this.getProducer() == nil {
		this.setProducer()
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
		Key:   sarama.StringEncoder(fmt.Sprintf("%d", time.Now().UnixNano())),
	}

	this.producer.Input() <- msg

	return nil
}
