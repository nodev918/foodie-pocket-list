package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		fmt.Println("get")
	})
	router.POST("/post", func(c *gin.Context) {
		fmt.Println("post")
	})

	router.Run(":8888")
}
