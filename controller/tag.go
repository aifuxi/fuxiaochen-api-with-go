package controller

import (
	"fmt"
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(ctx *gin.Context) {
	var tags []model.Tag

	res := global.DB.Find(&tags)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tags":   tags,
	})
}

func GetTag(ctx *gin.Context) {
	id := ctx.Param("id")
	var tag model.Tag

	res := global.DB.First(&tag, id)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tag":    tag,
	})
}

type createTagParams struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon"`
}

func CreateTag(ctx *gin.Context) {
	var params createTagParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	tag := model.Tag{
		Name: params.Name,
		Icon: params.Icon,
	}

	res := global.DB.Create(&tag)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"tag":    tag,
	})
}

type updateTagParams struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func UpdateTag(ctx *gin.Context) {
	id := ctx.Param("id")
	var tag model.Tag
	var params updateTagParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	res := global.DB.First(&tag, id)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	res = global.DB.Model(&tag).Updates(&model.Tag{
		Name: params.Name,
		Icon: params.Icon,
	})
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     fmt.Sprintf("%d", tag.ID),
	})
}

func DeleteTag(ctx *gin.Context) {
	id := ctx.Param("id")
	var tag model.Tag

	res := global.DB.First(&tag, id)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	res = global.DB.Delete(&model.Tag{}, id)
	if res.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "fail",
			"message": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     fmt.Sprintf("%d", tag.ID),
	})
}
