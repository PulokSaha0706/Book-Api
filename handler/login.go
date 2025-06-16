package handler

import (
	"BookApi/data"
	"BookApi/middleware"
	"BookApi/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for _, u := range data.Users {
		if u.Email == newUser.Email {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(newUser.Password))
			if err == nil {
				claims := &middleware.Claims{
					Username: newUser.Email,
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, err := token.SignedString(middleware.JwtKey)

				//fmt.Println(tokenString)

				if err != nil {
					c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not create token"})
					return
				}
				c.SetCookie("token", tokenString, 300, "/", "localhost", false, true)
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully"})

				return
			}

		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid credential"})

}
