package gooServer

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	DoHandle(c *gin.Context)
}
