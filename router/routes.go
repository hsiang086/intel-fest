package router

import (
  // "fmt"

	"github.com/gin-gonic/gin"
	"github.com/hsiang086/intel-fest/router/api"
)

func Root(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "sucess"})
}

func Routes() {
	Router.GET("/", Root)

	apiRoutes := Router.Group("/api")
	{
		apiRoutes.POST("/login", api.Login)
	}
}
