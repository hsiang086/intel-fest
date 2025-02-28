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

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func generateId() int {
	return rand.Intn(10000000)
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
		setCookie(c, content.Email, id)
		c.JSON(http.StatusOK, gin.H{"msg": "User created", "id": res, "success": true})
	}
}

func Login(c *gin.Context) {
	var content UserLogin
	if err := c.ShouldBindJSON(&content); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userExist, userId := database.IsUserExist(content.Email)
	if !userExist {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}
	user, email := database.GetUser(userId)
	if token, err := c.Cookie("__yumm__"); err == nil {
		if isCookieValid, _ := IsCookieValid(token); !isCookieValid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
			return
		}
		setCookie(c, content.Email, userId)
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("Login success as %s (%s)", user, email), "success": true})
	} else if database.IsUserPasswordValid(content.Email, content.Password) {
		setCookie(c, content.Email, userId)
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("Login success as %s (%s)", user, email), "success": true})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
	}
}
