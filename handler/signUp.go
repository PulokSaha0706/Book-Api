package handler

import (
	"BookApi/data"
	"BookApi/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for _, u := range data.Users {
		if u.Email == newUser.Email {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User alrerady exists"})
			return
		}
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Problem to create hash"})
		return
	}

	newUser.Password = string(hashedPass)
	data.Users = append(data.Users, newUser)

	c.IndentedJSON(http.StatusOK, data.Users)

}
