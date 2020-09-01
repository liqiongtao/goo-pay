package main

import (
	"context"
	"fmt"
	"googo.io/goo-mq"
	gooLog "googo.io/goo/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	topic = "test"
	addrs = []string{"s100:9092"}

	wg sync.WaitGroup

	sig         = make(chan os.Signal)
	ctx, cancel = context.WithCancel(context.Background())
)

func init() {
	gooMQ.Init(&gooMQ.Kafka{
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
	go publish()
	wg.Add(1)
	go publish()
	wg.Add(1)
	go publish()

	wg.Add(1)
	go consume()
	wg.Add(1)
	go consume()
	wg.Add(1)
	go consume()

	wg.Add(1)
	go consumeGroup("A100")
	wg.Add(1)
	go consumeGroup("A100")
	wg.Add(1)
	go consumeGroup("A101")
	wg.Add(1)
	go consumeGroup("A101")

	wg.Wait()
}

func publish() {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		msg := []byte(fmt.Sprintf("msg-%d", i))
		if err := gooMQ.SendMessage(topic, msg); err != nil {
			gooLog.Error(err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func consume() {
	defer wg.Done()

	ch, err := gooMQ.Consume(topic)
	if err != nil {
		gooLog.Error(err.Error())
		return
	}

	for {
		select {
		case buf := <-ch:
			gooLog.Debug(">>aa>>", string(buf))
		case <-ctx.Done():
			return
		}
	}
}

func consumeGroup(groupId string) {
	defer wg.Done()

	ch, err := gooMQ.ConsumeGroup(groupId, []string{topic})
	if err != nil {
		return
	}

	for {
		select {
		case buf := <-ch:
			gooLog.Debug(">>"+groupId+">>", string(buf))
		case <-ctx.Done():
			return
		}
	}
}
