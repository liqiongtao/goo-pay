package gooMQ

import (
	"fmt"
	"github.com/Shopify/sarama"
	gooLog "googo.io/goo/log"
	"sync"
)

type KafkaConsumer struct {
	*Kafka
	Topic  string
	Output chan []byte
	sync.WaitGroup
}

func (this *KafkaConsumer) getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_10_2_0
	return config
}

func (this *KafkaConsumer) Consume() error {
	consumer, err := sarama.NewConsumer(this.Addrs, this.getConfig())
	if err != nil {
		gooLog.Error("[kafka-consumer-error]", err.Error())
		return err
	}

	go func() {
		defer consumer.Close()

		partitions, err := consumer.Partitions(this.Topic)
		if err != nil {
			gooLog.Error("[kafka-consumer-error]", err.Error())
			return
		}

		for _, partition := range partitions {
			pc, err := consumer.ConsumePartition(this.Topic, partition, sarama.OffsetNewest)
			if err != nil {
				gooLog.Error("[kafka-consumer-error]", err.Error())
				continue
			}
			this.WaitGroup.Add(1)
			go this.ConsumePartition(pc)
		}

		this.WaitGroup.Wait()
	}()

	return nil
}

func (this *KafkaConsumer) ConsumePartition(pc sarama.PartitionConsumer) {
	defer this.WaitGroup.Done()
	defer pc.Close()

	for {
		select {
		case msg := <-pc.Messages():
			this.Output <- msg.Value
			gooLog.Debug("[kafka-consume-succ]", fmt.Sprintf("partitions=%d topic=%s offset=%d key=%s value=%s",
				msg.Partition, msg.Topic, msg.Offset, string(msg.Key), string(msg.Value)))
		case err := <-pc.Errors():
			gooLog.Error("[kafka-consumer-error]", err.Error())
		case <-this.Context.Done():
			return
		}
	}
}
