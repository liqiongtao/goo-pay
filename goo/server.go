package goo

import (
	"github.com/gin-gonic/gin"
	"goo/log"
	"goo/server"
)

func Gin() *gooServer.GinEngine {
	return gooServer.NewGin()
}

func Handler(controller gooServer.IController) gin.HandlerFunc {
	return controller.DoHandle
}

func Exception(code int, message string) {
	panic(gooServer.Response{
		Code:    code,
		Message: message,
		Data:    map[string]string{},
	})
}

func Success(data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	panic(gooServer.Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

func AsyncFunc(f func()) {
	go func(f func()) {
		defer func() {
			if err := recover(); err != nil {
				gooLog.Error(err)
			}
		}()

		f()
	}(f)
}
