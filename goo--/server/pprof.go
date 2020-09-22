package gooServer

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func PPROFRegister(r *gin.Engine, baseUri string) {
	pprof.Register(r, baseUri)
}

func PPROFRouteRegister(rg *gin.RouterGroup) {
	pprof.RouteRegister(rg)
}
