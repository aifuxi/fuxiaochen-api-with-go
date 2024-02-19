package middleware

import (
	"fuxiaochen-api-with-go/controller"
	"fuxiaochen-api-with-go/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			controller.ResponseError(c, controller.CodeTokenNotFound)
			c.Abort()
			return
		}

		parts := strings.SplitN(bearerToken, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			controller.ResponseError(c, controller.CodeIncorrectTokenFormat)
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := util.ParseToken(token)
		if err != nil {
			controller.ResponseError(c, controller.CodeIncorrectToken)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
