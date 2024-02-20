package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIDFromParam(c *gin.Context, id *int64) (err error) {
	tmp := c.Param("id")
	var tmpID int64

	if tmpID, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		return err
	}
	id = &tmpID
	return nil
}
