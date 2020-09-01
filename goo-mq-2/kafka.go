package gooMQ

import "context"

type Kafka struct {
	Addrs []string
	context.Context
}

func (this *Kafka) Init() {
}

func (this *Kafka) SendMessage(topic string, value []byte) error {
	c := &KafkaConsumer{
		Kafka:  this,
		Topic:  topic,
		Output: data,
	}
	return c.Consume()
}

func (this *Kafka) Consume(topic string, data chan []byte) error {
	c := &KafkaConsumer{
		Kafka:  this,
		Topic:  topic,
		Output: data,
	}
	return c.Consume()
}

func (this *Kafka) ConsumeGroup(groupId string, topics []string, data chan []byte) error {
	cg := &KafkaConsumerGroup{
		Kafka:   this,
		Topics:  topics,
		GroupId: groupId,
		Output:  data,
	}
	return cg.Consume()
}
