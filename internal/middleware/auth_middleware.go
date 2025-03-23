package middleware

import (
	"fastbuy/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware Middleware to check JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, status := auth.ParseToken(tokenString)
		if status != auth.TOKEN_VALID {
			c.JSON(http.StatusUnauthorized, gin.H{"status": status, "error": "Invalid token"})
			c.Abort()
			return
		}

		// Save username to context for use in other APIs
		c.Set("username", claims["username"])
		c.Next()
	}
}
