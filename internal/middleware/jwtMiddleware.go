package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token not provided"})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set("tokenString", tokenString)
	c.Set("token", token)
	c.Next()
}
