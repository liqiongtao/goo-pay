package gooMQ

import (
	"fmt"
	"github.com/Shopify/sarama"
	gooLog "googo.io/goo/log"
)

type KafkaConsumerGroup struct {
	*Kafka
	GroupId string
	Topics  []string
	Output  chan []byte
}

func (this *KafkaConsumerGroup) getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V0_10_2_0
	return config
}

func (this *KafkaConsumerGroup) Consume() error {
	cg, err := sarama.NewConsumerGroup(this.Addrs, this.GroupId, this.getConfig())
	if err != nil {
		gooLog.Error("[kafka-consumer-group-error]", err.Error())
		return err
	}

	go func() {
		defer cg.Close()
		for {
			if err := cg.Consume(this.Context, this.Topics, this); err != nil {
				gooLog.Error("[kafka-consumer-group-error]", err.Error())
				continue
			}
			if err := this.Context.Err(); err != nil {
				break
			}
		}
	}()

	return nil
}

func (this *KafkaConsumerGroup) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (this *KafkaConsumerGroup) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (this *KafkaConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		this.Output <- msg.Value

		// 更新位移
		sess.MarkMessage(msg, "")

		gooLog.Debug("[kafka-consume-group-succ]", fmt.Sprintf("partitions=%d topic=%s offset=%d key=%s groupid=%s value=%s",
			msg.Partition, msg.Topic, msg.Offset, string(msg.Key), sess.MemberID(), string(msg.Value)))
	}

	return nil
}
