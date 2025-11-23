package middleware

import (
	"net/http"
	"strings"

	"gin/internal/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(t *token.TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenStr := strings.Split(auth, "Bearer ")[1]
		parsed, err := t.VerifyAccessToken(tokenStr)
		if err != nil || !parsed.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims := parsed.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64)) //? what does this mean

		c.Set("user_id", userID)

		c.Next()
	}
}
