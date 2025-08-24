package middleware

import (
	"encoding/base64"
	"net/http"
	"quiz/config"
	"quiz/structs"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorizaton Required"})
			return
		}
		code, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		parts := strings.SplitN(string(code), ":", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			return
		}
		username, password := parts[0], parts[1]

		var user structs.User
		if err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		c.Next()
	}

}
