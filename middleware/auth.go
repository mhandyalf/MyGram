package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginAuth(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil || !IsTokenValid(sessionToken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.Next()
}

func IsTokenValid(token string) bool {
	return token != ""
}
