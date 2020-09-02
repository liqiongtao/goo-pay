package gooMQ_v2

import (
	"github.com/streadway/amqp"
	gooLog "googo.io/goo/log"
	"sync"
)

type Rabbitmq struct {
	Url  string
	conn *amqp.Connection
	ch   *amqp.Channel
	mu   sync.Mutex
}

func (this *Rabbitmq) Connect() (err error) {
	this.conn, err = amqp.Dial(this.Url)
	if err != nil {
		gooLog.Error(err.Error())
		return
	}
	this.ch, err = this.conn.Channel()
	if err != nil {
		gooLog.Error(err.Error())
	}
	return
}

func (this *Rabbitmq) Publish(topic string, body []byte) error {
	q, err := this.getChannel().QueueDeclare(topic, false, false, false, false, nil)
	if err != nil {
		gooLog.Error(err.Error())
		return err
	}
	return this.getChannel().Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

func (this *Rabbitmq) Consume(topic string, data chan []byte) error {
	q, err := this.getChannel().QueueDeclare(topic, false, false, false, false, nil)
	if err != nil {
		gooLog.Error(err.Error())
		return err
	}

	msgs, err := this.getChannel().Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		gooLog.Error(err.Error())
		return err
	}

	go func() {
		for msg := range msgs {
			data <- msg.Body
		}
	}()

	return nil
}

func (this *Rabbitmq) getChannel() *amqp.Channel {
	if this.ch == nil {
		this.Connect()
	}
	return this.ch
}
