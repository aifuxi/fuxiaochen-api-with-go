package router

import (
	"fuxiaochen-api-with-go/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
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
		adminTagRouter.GET("/", controller.GetTags)
		adminTagRouter.GET("/:id", controller.GetTag)
		adminTagRouter.POST("/", controller.CreateTag)
		adminTagRouter.PATCH("/:id", controller.UpdateTag)
		adminTagRouter.DELETE("/:id", controller.DeleteTag)
	}

	return r
}
