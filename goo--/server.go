package goo__

import (
	"github.com/gin-gonic/gin"
	gooLog "github.com/liqiongtao/googo.io/goo/log"
	gooServer "github.com/liqiongtao/googo.io/goo/server"
)

func Gin() *gooServer.GinEngine {
	return gooServer.NewGin()
}

func Handler(controller gooServer.IController) gin.HandlerFunc {
	return controller.DoHandle
}

func Exception(code int, message string) {
	panic(ExceptionJSON(code, message))
}

func ExceptionJSON(code int, message string) gooServer.Response {
	return gooServer.Response{
		Status:  0,
		Code:    code,
		Message: message,
		Data:    map[string]string{},
	}
}

func Success(data interface{}) {
	panic(SuccessJSON(data))
}

func SuccessJSON(data interface{}) gooServer.Response {
	if data == nil {
		data = map[string]string{}
	}
	return gooServer.Response{
		Status:  1,
		Code:    200,
		Message: "ok",
		Data:    data,
	}
}

func AsyncFunc(fn func()) {
	go func(fn func()) {
		defer func() {
			if err := recover(); err != nil {
				gooLog.Error(err)
			}
		}()
		fn()
	}(fn)
}
