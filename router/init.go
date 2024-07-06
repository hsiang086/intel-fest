package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Host   string
	Port   int
	Router *gin.Engine
	Server *http.Server
)

func Init() {
	Host = "http://127.0.0.1"
	Port = 1588
	Router = gin.Default()
	Server = &http.Server{
		Addr:           fmt.Sprintf(":%d", Port),
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	Router.Use(func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})
	fmt.Printf("Initialized starting at: %s:%d\n", Host, Port)
}
