package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Engine
	router := gin.Default()
	//router := gin.New()
	asd := 122
	log.Printf("%d", asd)
	router.GET("/hello", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200, gin.H{
			"code":    200,
			"success": true,
		})
	})
	// 指定地址和端口号
	router.Run("localhost:9090")
}
