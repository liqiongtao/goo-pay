package gooMQ

type imq interface {
	Init()
	SendMessage(topic string, value []byte) error
	Consume(topic string, data chan []byte) error
	ConsumeGroup(groupId string, topics []string, data chan []byte) error
}

var __mq imq

func Init(mq imq) {
	__mq = mq
	__mq.Init()
}

func SendMessage(topic string, value []byte) error {
	return __mq.SendMessage(topic, value)
}

func Consume(topic string) (<-chan []byte, error) {
	data := make(chan []byte)
	return (<-chan []byte)(data), __mq.Consume(topic, data)
}

func ConsumeGroup(groupId string, topics []string) (<-chan []byte, error) {
	data := make(chan []byte)
	return (<-chan []byte)(data), __mq.ConsumeGroup(groupId, topics, data)
}
