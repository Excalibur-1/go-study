package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pang",
		})
	})
	// 在 0.0.0.0:8080 上监听并服务
	err := engine.Run()
	if err != nil {
		log.Fatalf("engine.Run() err : %s", err)
	}
}
