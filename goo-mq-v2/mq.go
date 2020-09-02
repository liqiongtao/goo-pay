package gooMQ_v2

type HandlerFunc func(data []byte) bool

type imq interface {
	Init()
	SendMessage(topic string, value []byte) error
	Consume(topic string, handler HandlerFunc) error
	ConsumeGroup(groupId string, topics []string, handler HandlerFunc) error
}

var __mq imq

func Init(mq imq) {
	__mq = mq
	__mq.Init()
}

func SendMessage(topic string, value []byte) error {
	return __mq.SendMessage(topic, value)
}

func Consume(topic string, handler HandlerFunc) error {
	return __mq.Consume(topic, handler)
}

func ConsumeGroup(groupId string, topics []string, handler HandlerFunc) error {
	return __mq.ConsumeGroup(groupId, topics, handler)
}
