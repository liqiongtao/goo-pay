package main

import (
	"context"
	"fmt"
	gooMQ "github.com/liqiongtao/googo.io/goo-mq"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	topic  = "test"
	topics = []string{"test"}
	addrs  = []string{"s100:9092"}

	sig         = make(chan os.Signal)
	ctx, cancel = context.WithCancel(context.Background())
)

func init() {
	gooMQ.Init(&gooMQ.Kafka{Context: ctx, Addrs: addrs})
	// gooMQ.Init(&gooMQ.KafkaProducer{})
	// gooMQ.Init(&gooMQ.KafkaConsumer{})
	// gooMQ.Init(&gooMQ.KafkaConsumerGroup{})

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for ch := range sig {
			switch ch {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				cancel()
			default:
				continue
			}
		}
	}()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			gooMQ.SendMessage(topic, []byte(fmt.Sprintf("msg-%d", i)))
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		gooMQ.Consume(topic, func(data []byte) bool {
			return true
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gooMQ.Consume(topic, func(data []byte) bool {
			return true
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gooMQ.ConsumeGroup("A100", topics, func(data []byte) bool {
			return true
		})
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		gooMQ.ConsumeGroup("A101", topics, func(data []byte) bool {
			return true
		})
	}()

	wg.Wait()
}
