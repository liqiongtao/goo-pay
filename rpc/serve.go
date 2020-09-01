package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	rpc_wx "wx.googo.io/protos.pb"
)

func Serve(svr rpc_wx.WXServiceServer) {
	pid := fmt.Sprintf("%d", os.Getpid())
	if err := ioutil.WriteFile(".pid", []byte(pid), 0755); err != nil {
		panic(err.Error())
	}

	lis, err := net.Listen("tcp", Config.Server.Port)
	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()

	rpc_wx.RegisterWXServiceServer(s, svr)
	reflection.Register(s)

	go func() {
		log.Printf("serve running on %s", Config.Server.Port)
		if err := s.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.GracefulStop()
			return
		}
	}
}
