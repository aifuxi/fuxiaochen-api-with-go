package middleware

import (
	"fuxiaochen-api-with-go/constant"
	"fuxiaochen-api-with-go/controller"
	"fuxiaochen-api-with-go/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get(constant.Authorization)
		if bearerToken == "" {
			controller.ResponseError(c, controller.CodeTokenNotFound, nil)
			c.Abort()
			return
		}

		parts := strings.SplitN(bearerToken, " ", 2)
		if len(parts) != 2 || parts[0] != constant.Bearer {
			controller.ResponseError(c, controller.CodeIncorrectTokenFormat, nil)
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := util.ParseToken(token)
		if err != nil {
			controller.ResponseError(c, controller.CodeIncorrectToken, nil)
			c.Abort()
			return
		}

		c.Set(constant.UserID, claims.UserID)
		c.Next()
	}
}
