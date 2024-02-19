package router

import (
	"fuxiaochen-api-with-go/controller"
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/middleware"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(ginzap.Ginzap(global.Logger, time.DateTime, true))

	r.Use(ginzap.RecoveryWithZap(global.Logger, true))

	adminAPIV1 := r.Group("/admin-api/v1")
	authRouter := adminAPIV1.Group("/auth")
	{
		authRouter.POST("/login", controller.Login)
	}

	adminUserRouter := adminAPIV1.Group("/users", middleware.AuthMiddleware())
	{
		adminUserRouter.GET("", controller.GetUsers)
		adminUserRouter.POST("", controller.CreatUser)
		adminUserRouter.GET(":id", controller.GetUser)
		adminUserRouter.PATCH(":id", controller.UpdateUser)
		adminUserRouter.DELETE(":id", controller.DeleteUser)
	}

	adminTagRouter := adminAPIV1.Group("/tags", middleware.AuthMiddleware())
	{
		adminTagRouter.GET("", controller.GetTags)
		adminTagRouter.POST("", controller.CreatTag)
		adminTagRouter.GET(":id", controller.GetTag)
		adminTagRouter.PATCH(":id", controller.UpdateTag)
		adminTagRouter.DELETE(":id", controller.DeleteTag)
	}

	return r
}
