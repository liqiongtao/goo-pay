package gooMQ_v2

import (
	"fmt"
	"github.com/Shopify/sarama"
	gooLog "googo.io/goo/log"
	"sync"
)

type KafkaConsumer struct {
	*Kafka
	Topic   string
	Handler HandlerFunc
	wg      sync.WaitGroup
}

func (this *KafkaConsumer) config() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_10_2_0
	return config
}

func (this *KafkaConsumer) Consume() error {
	consumer, err := sarama.NewConsumer(this.Addrs, this.config())
	if err != nil {
		gooLog.Error("[kafka-consumer-error]", err.Error())
		panic(err.Error())
	}

	partitions, err := consumer.Partitions(this.Topic)
	if err != nil {
		gooLog.Error("[kafka-consumer-error]", err.Error())
		return err
	}

	go func(partitions []int32) {
		defer consumer.Close()

		for _, partition := range partitions {
			pc, err := consumer.ConsumePartition(this.Topic, partition, sarama.OffsetNewest)
			if err != nil {
				gooLog.Error("[kafka-consumer-error]", err.Error())
				continue
			}
			this.wg.Add(1)
			go this.message(pc)
		}

		this.wg.Wait()
	}(partitions)

	return nil
}

func (this *KafkaConsumer) message(pc sarama.PartitionConsumer) {
	defer this.wg.Done()
	defer pc.Close()

	for {
		select {
		case msg := <-pc.Messages():
			this.Handler(msg.Value)
			gooLog.Debug("[kafka-consume-succ]",
				fmt.Sprintf("partitions=%d topic=%s offset=%d key=%s value=%s",
					msg.Partition, msg.Topic, msg.Offset, string(msg.Key), string(msg.Value)))

		case err := <-pc.Errors():
			gooLog.Error("[kafka-consumer-error]", err.Error())

		case <-this.Context.Done():
			return
		}
	}
}
