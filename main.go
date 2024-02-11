package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run("127.0.0.1:6122")
	if err != nil {
		log.Fatalf("启动服务器失败：%v", err)
	}
}
