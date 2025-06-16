package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var JwtKey = []byte("fgdfg")

func AuthMiddleware(c *gin.Context) {

	fmt.Println("Middle")
	var tokenString string

	cookie, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
		return
	}
	tokenString = cookie

	fmt.Println(cookie)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return
	}

	// Store username in context for use in routes
	c.Set("username", claims.Username)
	c.Next()
}
