package gooServer

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type GinEngine struct {
	*gin.Engine
	noLogPaths map[string]struct{}
	requestId  int64
}

func NewGin() *GinEngine {
	g := &GinEngine{
		Engine:     gin.New(),
		noLogPaths: map[string]struct{}{},
		requestId:  0,
	}
	g.Use(cors(), noAccess(), logger(g), recovery())
	g.NoRoute(noRoute())
	return g
}

func (g *GinEngine) Serve(addr string) {
	pid := fmt.Sprintf("%d", os.Getpid())
	if err := ioutil.WriteFile(".pid", []byte(pid), 0755); err != nil {
		panic(err.Error())
	}
	endless.NewServer(addr, g.Engine).ListenAndServe()
}

func (g *GinEngine) SetNoLogPath(paths ...string) {
	for _, i := range paths {
		g.noLogPaths[i] = struct{}{}
	}
}
