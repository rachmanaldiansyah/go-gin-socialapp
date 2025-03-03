package middlewares

import (
	"go-gin-sosmed/exceptions"
	"go-gin-sosmed/helpers"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			exceptions.HandleError(c, &exceptions.UnauthorizedError{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helpers.ValidateToken(tokenString)
		if err != nil {
			exceptions.HandleError(c, &exceptions.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userId", *userId)
		c.Next()
	}
}