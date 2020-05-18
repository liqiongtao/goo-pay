package gooMQ

type Producer interface {
	Message() string
}

type Receiver interface {
	Consumer([]byte) error
}
