package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIDFromParam(c *gin.Context) (id int64, err error) {
	tmp := c.Param("id")

	if id, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		return id, err
	}

	return id, nil
}
