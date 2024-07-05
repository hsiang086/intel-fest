package api

import (
	"math/rand"
	"net/http"
	"strconv"

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
	userExist, userId := database.IsUserExist(content.Email)
	if userExist {
		c.JSON(http.StatusOK, gin.H{"msg": "User already exist", "id": userId})
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
		c.SetCookie("__yumm__", strconv.Itoa(id), 3600, "/", "http://127.0.0.1", false, true)
		c.JSON(http.StatusOK, gin.H{"msg": "User created", "id": res})
	}
}
