package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hsiang086/intel-fest/router/api"
)

func Root(c *gin.Context) {
	var msg gin.H
	if token, err := c.Cookie("__yumm__"); err == nil {
		if isCookieValid, email := api.IsCookieValid(token); isCookieValid {
			msg = gin.H{"msg": fmt.Sprintf("Welcome (%s)", email)}
		} else {
			msg = gin.H{"msg": "Please login first"}
		}
		c.JSON(200, msg)
	}
	c.Redirect(302, "/login")
}

func Login(c *gin.Context) {
	c.HTML(200, "auth/index.tmpl", gin.H{
		"use": "login",
	})
}

func Signup(c *gin.Context) {
	c.HTML(200, "auth/index.tmpl", gin.H{
		"use": "signup",
	})
}

func Routes() {
	Router.GET("/", Root)
	Router.GET("/login", Login)
	Router.GET("/signup", Signup)

	apiRoutes := Router.Group("/api")
	{
		apiRoutes.POST("/signup", api.Signup)
		apiRoutes.POST("/login", api.Login)
	}
}
