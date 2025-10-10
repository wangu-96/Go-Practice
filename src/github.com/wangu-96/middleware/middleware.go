package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/JWT"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Expect "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid Authorization header"})
			c.Abort()
			return
		}

		token := parts[1]
		if err := JWT.VerifyToken(token); err != nil {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Token is valid, continue
		c.Next()
	}
}
