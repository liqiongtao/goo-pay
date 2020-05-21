package main

import (
	"github.com/gin-gonic/gin"
	"googo.io/goo"
	gooCache "googo.io/goo/cache"
	gooConfig "googo.io/goo/config"
	gooDB "googo.io/goo/db"
)

var conf = &config{}

func init() {
	gooConfig.LoadFile(".yaml", conf)
	gooDB.Init(conf.Mysql)
	gooCache.Init(conf.Redis)
}

func main() {
	g := goo.Gin()

	app := g.Group("/app")
	app.Use(authorize())
	{
		app.GET("ping", goo.Handler(Ping{}))
		app.GET("user/info", goo.Handler(UserInfo{}))
	}

	g.Serve(conf.Server.Port)
}

func authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

type Ping struct {
}

func (p Ping) DoHandle(c *gin.Context) {
	c.String(200, "ping")
}

type UserInfo struct {
}

type User struct {
	Avatar string
}

func (u UserInfo) DoHandle(c *gin.Context) {
	rsp := map[string]interface{}{
		"name": "hnatao",
	}
	goo.Success(rsp)
}
