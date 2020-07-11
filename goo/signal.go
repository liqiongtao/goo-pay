package goo

import (
	"os"
	"os/signal"
	"syscall"
)

var sig *gooSignal

func init() {
	sig = &gooSignal{
		SCH:      make(chan os.Signal),
		IsExit:   false,
		IsExitCH: make(chan bool),
	}
	signal.Notify(sig.SCH, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
}

type gooSignal struct {
	SCH      chan os.Signal
	IsExit   bool
	IsExitCH chan bool
}

func (s *gooSignal) Listening() {
	AsyncFunc(func() {
		for sig := range s.SCH {
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				s.IsExitCH <- true
			case syscall.SIGUSR1:
			case syscall.SIGUSR2:
			default:
			}
		}

		s.IsExit = <-s.IsExitCH
	})
}

func IsExit() bool {
	return sig.IsExit
}
