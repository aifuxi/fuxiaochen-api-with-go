package controller

import (
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(ctx *gin.Context) {
	var users []model.User

	res := global.DB.Find(&users)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  users,
	})
}

type createUserParams struct {
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required,eqfield=ConfirmPassword,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
}

func CreatUser(ctx *gin.Context) {
	var params createUserParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	user := model.User{
		Name:     params.Name,
		Password: params.Password,
	}
	res := global.DB.Create(&user)

	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
}
