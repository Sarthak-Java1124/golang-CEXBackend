package middleware

import (
	"golang-CEX/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing authorization header",
				},
			)
			return
		}

		// Validate Bearer format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid authorization format",
				},
			)
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		// Verify JWT
		claims, err := utils.VerifyToken(
			tokenString,
			"my-secret-token",
		)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid or expired token",
				},
			)
			return
		}

		// Inject values into gin context
		c.Set("userId", claims.UserId)
		c.Set("claims", claims)

		// Continue request chain
		c.Next()
	}
}
