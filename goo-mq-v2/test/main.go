package main

import (
	"context"
	"googo.io/goo-mq-v2"
	"googo.io/goo/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	topic  = "test"
	topics = []string{topic}
	addrs  = []string{"s100:9092"}

	wg sync.WaitGroup

	sig         = make(chan os.Signal)
	ctx, cancel = context.WithCancel(context.Background())
)

func init() {
	gooMQ_v2.Init(&gooMQ_v2.Kafka{
		Context: ctx,
		Addrs:   addrs,
	})

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
	wg.Add(1)
	go func() {
		defer wg.Done()
		gooMQ_v2.ConsumeGroup("A100", topics, consume(1))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gooMQ_v2.ConsumeGroup("A101", topics, consume(1))
	}()

	wg.Wait()
}

func consume(id int) gooMQ_v2.HandlerFunc {
	return func(data []byte) bool {
		gooLog.Debug(">>>>>>>", string(data))
		if (id%2 == 0) {
			return false
		}
		return true
	}
}
