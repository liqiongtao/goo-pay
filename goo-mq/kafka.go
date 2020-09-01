package gooMQ

import (
	"context"
)

type Kafka struct {
	Addrs []string
	context.Context
	producer *KafkaProducer
}

func (this *Kafka) Init() {
	this.producer = &KafkaProducer{
		Kafka: this,
	}
}

func (this *Kafka) SendMessage(topic string, value []byte) error {
	return this.producer.SendMessage(topic, value)
}

func (this *Kafka) Consume(topic string, data chan []byte) error {
	consumer := &KafkaConsumer{
		Kafka:  this,
		Topic:  topic,
		Output: data,
	}
	return consumer.Consume()
}

func (this *Kafka) ConsumeGroup(groupId string, topics []string, data chan []byte) error {
	consumerGroup := &KafkaConsumerGroup{
		Kafka:   this,
		GroupId: groupId,
		Topics:  topics,
		Output:  data,
	}
	return consumerGroup.Consume()
}
