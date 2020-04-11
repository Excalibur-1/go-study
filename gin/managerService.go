package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {

	})
	//使用endless替换默认的ListenAndServe
	err := endless.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
