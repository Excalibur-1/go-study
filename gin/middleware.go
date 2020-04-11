package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func MyBenchLogger() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := context.Request.Header.Get("user")
		password := context.Request.Header.Get("password")
		if user != "admin" || password != "123" {
			panic("user not found or password is error")
		}
	}
}

func main() {
	//写入日志文件
	//禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色
	gin.DisableConsoleColor()

	file, err := os.Create("gin.log")
	if err != nil {
		log.Fatalf("os.Create err: %s", err)
	}
	gin.DefaultWriter = io.MultiWriter(file)

	//如果你需要同时写入日志文件和控制台上显示，使用下面代码
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	//使用中间件
	//创建一个默认的没有任何中间件的路由
	r := gin.New()

	//全局中间件
	//Logger 中间件将日志写到gin.DefaultWriter，即使你设置GIN_MODE=release
	//默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	//Recovery 中间件从任何panic恢复，如果出现panic，它会写一个500错误
	r.Use(gin.Recovery())

	//每个路由的中间件，你能添加任意数量的中间件
	r.GET("/benchmark", MyBenchLogger(), func(context *gin.Context) {})

	//授权组
	//authorized := r.Group("/", AuthRequired())
	//也可以这样：
	authorized := r.Group("/")
	//每个组的中间价！在这个实例中，我们只需要在“authorized”组中
	//使用自定义创建的AuthRequired()中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", func(context *gin.Context) {})
		authorized.POST("/submit", func(context *gin.Context) {})
		authorized.POST("/read", func(context *gin.Context) {})

		//nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(context *gin.Context) {})
	}

	//默认情况下，它使用：8080，除非定义了PORT环境变量
	err = r.Run()
	//硬编码端口
	//err := router.Run(":3000")
	if err != nil {
		log.Fatalf("router.Run err: %s", err)
	}

}
