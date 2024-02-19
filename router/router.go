package router

import (
	"fuxiaochen-api-with-go/controller"
	"fuxiaochen-api-with-go/global"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.GinzapWithConfig(global.Logger, &ginzap.Config{
		UTC:       false,
		SkipPaths: nil,
	}))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(global.Logger, true))

	adminAPIV1 := r.Group("/admin-api/v1")
	adminUserRouter := adminAPIV1.Group("/users")
	{
		adminUserRouter.GET("", controller.GetUsers)
		adminUserRouter.POST("", controller.CreatUser)
		adminUserRouter.GET(":id", controller.GetUser)
		adminUserRouter.PATCH(":id", controller.UpdateUser)
		adminUserRouter.DELETE(":id", controller.DeleteUser)
	}

	adminTagRouter := adminAPIV1.Group("/tags")
	{
		adminTagRouter.GET("", controller.GetTags)
		adminTagRouter.POST("", controller.CreatTag)
		adminTagRouter.GET(":id", controller.GetTag)
		adminTagRouter.PATCH(":id", controller.UpdateTag)
		adminTagRouter.DELETE(":id", controller.DeleteTag)
	}

	return r
}
