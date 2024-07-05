package router

import (
	"fmt"
	"strconv"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Port = 1588
	PortStr = strconv.Itoa(Port)
	Router = gin.Default()
	Server = &http.Server{
		Addr:           ":" + PortStr,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
)

func Init() {
	Router.Use(func (c *gin.Context){
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})
	fmt.Printf("Initialized starting at: http://127.0.0.1:%d\n", Port)
}
