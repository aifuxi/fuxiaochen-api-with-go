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
        adminUserRouter.GET("/", controller.GetUsers)
        adminUserRouter.POST("/", controller.CreatUser)
    }

    return r
}
