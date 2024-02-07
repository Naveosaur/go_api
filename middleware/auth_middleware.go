package middleware

import (
	errorhandler "go_gin/errorHandler"
	"go_gin/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "Unathorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}
}