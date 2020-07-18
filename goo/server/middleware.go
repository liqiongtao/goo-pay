package gooServer

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	gooLog "googo.io/goo/log"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

func logger(g *GinEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		if _, ok := g.noLogPaths[path]; ok {
			return
		}

		g.requestId += 1

		body := ""
		switch c.ContentType() {
		case "application/x-www-form-urlencoded", "application/json", "text/xml":
			buf, _ := ioutil.ReadAll(c.Request.Body)
			body = string(buf)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		}

		var bf bytes.Buffer

		bf.WriteString("{")
		bf.WriteString(fmt.Sprintf("\"request_id\":%d,", g.requestId))
		bf.WriteString(fmt.Sprintf("\"auth_user_id\":%d,", c.GetInt64("authUserId")))
		bf.WriteString(fmt.Sprintf("\"method\":\"%s\",", c.Request.Method))
		bf.WriteString(fmt.Sprintf("\"uri\":\"%s\",", c.Request.RequestURI))
		bf.WriteString(fmt.Sprintf("\"body\":\"%s\",", body))
		bf.WriteString(fmt.Sprintf("\"authorization\":\"%s\",", c.GetHeader("Authorization")))
		bf.WriteString(fmt.Sprintf("\"x-request-id\":\"%s\",", c.GetHeader("X-Request-Id")))
		bf.WriteString(fmt.Sprintf("\"x-request-timestamp\":\"%s\",", c.GetHeader("X-Request-Timestamp")))
		bf.WriteString(fmt.Sprintf("\"x-request-sign\":\"%s\",", c.GetHeader("X-Request-Sign")))
		bf.WriteString(fmt.Sprintf("\"content-type\":\"%s\",", c.ContentType()))
		bf.WriteString(fmt.Sprintf("\"client_ip\":\"%s\",", c.ClientIP()))
		bf.WriteString(fmt.Sprintf("\"referer\":\"%s\",", c.GetHeader("Referer")))

		rsp, ok := c.Get("Response")
		if ok {
			bf.WriteString(fmt.Sprintf("\"response\":%s,", rsp.(Response).ToString()))
		}

		bf.WriteString(fmt.Sprintf("\"execution_time\":\"%dms\"", (time.Now().UnixNano()-start.UnixNano())/1e6))
		bf.WriteString("}")

		if ok && rsp.(Response).Code != 0 {
			gooLog.Error(bf.String())
		} else {
			gooLog.Debug(bf.String())
		}
	}
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch reflect.TypeOf(err).String() {
				case "gooServer.Response":
					c.Set("Response", err.(Response))
					c.JSON(http.StatusOK, err.(Response))
				default:
					c.String(http.StatusOK, fmt.Sprint(err))
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}

func cors() gin.HandlerFunc {
	allowHeaders := "Content-Type, Content-Length, Authorization, Accept, Referer, User-Agent, " +
		"X-Requested-Id, X-Request-Timestamp, X-Request-Sign, X-Request-AppId, X-Request-Source, X-Request-Token"

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", allowHeaders)
		c.Next()
	}
}

func noAccess() gin.HandlerFunc {
	noAccessPaths := []string{
		"/favicon.ico",
	}

	noAccessPathsMap := map[string]struct{}{}
	for _, i := range noAccessPaths {
		noAccessPathsMap[i] = struct{}{}
	}

	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		if _, ok := noAccessPathsMap[c.Request.URL.Path]; ok {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func noRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(Response{
			Code:    404,
			Message: "Page Not Found",
			Data:    map[string]string{},
		})
	}
}
