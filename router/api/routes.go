package api

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsiang086/intel-fest/database"
)

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func generateId() int {
	return rand.Intn(1000000)
}

func Signup(c *gin.Context) {
	var content User
	if err := c.ShouldBindJSON(&content); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if database.IsUserExist(content.Email) {
		c.JSON(http.StatusOK, gin.H{"msg": "User already exist"})
	} else {
		id := generateId()
		for {
			if database.IsUserIdExist(id) {
				id = generateId()
			} else {
				break
			}
		}
		res := database.InsertUser(id, content.Name, content.Email, content.Password)
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("User %d created", res)})
	}
}
