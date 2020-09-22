package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"googo.io/goo"
)

var (
	ctx = context.Background()
)

func main() {
	s := goo.NewServer()
	s.GET("/", goo.Handler(UserList{}))
	s.Run(":18000")
}

type UserList struct {
	Name string `form:"name"`
}

func (ul UserList) DoHandle(c *gin.Context) *goo.Response {
	if err := c.ShouldBind(&ul); err != nil {
		return goo.Error(40010, "参数错误", err.Error())
	}
	return goo.Success(gin.H{"name": ul.Name})
}
