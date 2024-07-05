package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id          int `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var content User
	if err := c.ShouldBindJSON(&content); err!= nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, content)
}
