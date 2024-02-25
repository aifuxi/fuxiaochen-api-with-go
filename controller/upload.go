package controller

import (
	"fmt"
	"fuxiaochen-api-with-go/global"
	"fuxiaochen-api-with-go/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Upload(c *gin.Context) {
	var file *multipart.FileHeader
	var filename string
	var exists bool
	var err error

	if file, err = c.FormFile("file"); err != nil {
		ResponseError(c, CodeInternalServerError, err)
	}
	if exists, err = util.PathExists(global.Conf.AppConfig.UploadDir); !exists || err != nil {
		if err = os.Mkdir(global.Conf.AppConfig.UploadDir, os.ModePerm); err != nil {
			ResponseError(c, CodeInternalServerError, err)
		}
	}

	fileID, _ := uuid.NewUUID()
	filename = fmt.Sprintf("%s-%s", fileID, filepath.Base(file.Filename))

	if err = c.SaveUploadedFile(file, filepath.Join(global.Conf.AppConfig.UploadDir, filename)); err != nil {
		ResponseError(c, CodeInternalServerError, err)
	}

	ResponseSuccess(c, gin.H{
		"url": fmt.Sprintf(
			"%s:%d/%s/%s",
			global.Conf.AppConfig.Host,
			global.Conf.AppConfig.Port,
			global.Conf.AppConfig.StaticPath, filename),
	})
}
