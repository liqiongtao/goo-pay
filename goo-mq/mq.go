package gooMQ

// func MQ() *mq {
// 	return &mq{}
// }

// type mq struct {
// 	receiverList []Receiver
// }

// func (this *mq) RegisterReceiver(receiver Receiver) {
// 	this.receiverList = append(this.receiverList, receiver)
// }

// func (this *mq) Start() {
// 	for _, receiver := range this.receiverList {
// 		go this.listenReceiver(receiver)
// 	}
// }

// func (this *mq) listenReceiver(receiver Receiver) {

// }

type imq interface {
	Connect() error
	Publish(topic string, body []byte) error
	Consume(topic string, data chan []byte) error
}

var __mq imq

func Init(mq imq) {
	__mq = mq
	__mq.Connect()
}

func Publish(topic string, body []byte) error {
	return __mq.Publish(topic, body)
}

func Consume(topic string) (<-chan []byte, error) {
	data := make(chan []byte)
	return (<-chan []byte)(data), __mq.Consume(topic, data)
}
