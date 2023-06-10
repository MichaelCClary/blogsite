package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcclary/blogsite/helper"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentiction required"})
			context.Abort()
			return
		}
		context.Next()
	}
}
