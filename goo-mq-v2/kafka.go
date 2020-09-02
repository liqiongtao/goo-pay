package gooMQ_v2

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

func (this *Kafka) Consume(topic string, handler HandlerFunc) error {
	consumer := &KafkaConsumer{
		Kafka:   this,
		Topic:   topic,
		Handler: handler,
	}
	return consumer.Consume()
}

func (this *Kafka) ConsumeGroup(groupId string, topics []string, handler HandlerFunc) error {
	consumerGroup := &KafkaConsumerGroup{
		Kafka:   this,
		GroupId: groupId,
		Topics:  topics,
		Handler: handler,
	}
	return consumerGroup.Consume()
}
