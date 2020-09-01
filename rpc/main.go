package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"googo.io/rpc/protos.pb"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	isQuit bool
)

func main() {
	wpid()

	lis, err := net.Listen("tcp", ":13009")
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()

	rpc_goo.RegisterCaptchaServiceServer(s, &CaptchaService{})

	go func() {
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			panic(err.Error())
		}
	}()

	signalMonitor(s)
}

func wpid() {
	pid := fmt.Sprintf("%d", os.Getpid())
	if err := ioutil.WriteFile(".pid", []byte(pid), 0755); err != nil {
		panic(err.Error())
	}
}

func signalMonitor(s *grpc.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for {
			switch <-ch {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				isQuit = true
				s.GracefulStop()
				return
			}
		}
	}()
}
