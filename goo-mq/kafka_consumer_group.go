package gooMQ

import (
	"fmt"
	"github.com/Shopify/sarama"
	gooLog "github.com/liqiongtao/googo.io/goo/log"
)

type KafkaConsumerGroup struct {
	*Kafka
	GroupId string
	Handler HandlerFunc
}

func (*KafkaConsumerGroup) config() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V0_10_2_0
	return config
}

func (cg *KafkaConsumerGroup) Init() {
	gooLog.Debug("[kafka-consumer-group-init]")
}

func (cg *KafkaConsumerGroup) Consume(topics []string, handler HandlerFunc) error {
	c, err := sarama.NewConsumerGroup(cg.Addrs, cg.GroupId, cg.config())
	if err != nil {
		gooLog.Error("[kafka-consumer-group-error]", err.Error())
		return err
	}
	defer c.Close()

	cg.Handler = handler

	for {
		if err := c.Consume(cg.Context, topics, cg); err != nil {
			gooLog.Error("[kafka-consumer-group-error]", err.Error())
			continue
		}
		if err := cg.Context.Err(); err != nil {
			break
		}
	}

	return nil
}

func (cg *KafkaConsumerGroup) Setup(sess sarama.ConsumerGroupSession) (err error) {
	return
}

func (cg *KafkaConsumerGroup) Cleanup(sess sarama.ConsumerGroupSession) (err error) {
	return
}

func (cg *KafkaConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) (err error) {
	for msg := range claim.Messages() {
		flag := cg.Handler(msg.Value)

		if !flag {
			// 重置位移
			sess.ResetOffset(msg.Topic, msg.Partition, msg.Offset, "")
			return
		}

		// 更新位移
		sess.MarkMessage(msg, "")

		gooLog.Debug("[kafka-consumer-group-success]",
			fmt.Sprintf("partitions=%d topic=%s offset=%d key=%s groupid=%s value=%s",
				msg.Partition, msg.Topic, msg.Offset-1, string(msg.Key), cg.GroupId, string(msg.Value)))
	}

	return
}
