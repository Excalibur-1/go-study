package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	//使用默认中间件创建一个gin路由
	//日志与恢复中间件（无崩溃 ）
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {})
	router.POST("/somePost", func(c *gin.Context) {})
	router.PUT("/somePut", func(c *gin.Context) {})
	router.DELETE("/someDelete", func(c *gin.Context) {})
	router.PATCH("/somePatch", func(c *gin.Context) {})
	router.HEAD("/someHead", func(c *gin.Context) {})
	router.OPTIONS("/someOptions", func(c *gin.Context) {})

	//path中的参数（路径参数）
	//这个处理器可以匹配 /user/john，但是它不会匹配 /user
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	//但是，这个可以匹配 /user/john 和 /user/john/send
	//如果没有其他的路由匹配/user/john，他将重定向到/user/john/
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	//查询字符串参数
	//查询字符串参数使用现有的底层 request 对象解析
	//请求响应匹配的URL：/welcome?firstName=Jane&lastName=Doe
	router.GET("/welcome", func(context *gin.Context) {
		firstName := context.DefaultQuery("firstName", "Guest")
		//lastName := context.Request.URL.Query().Get("lastName")
		//这个是 context.Request.URL.Query().Get("lastName") 的快捷方式
		lastName := context.Query("lastName")

		context.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	//Multipart/Urlencoded 表单
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//查询字符串 + post表单
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		log.Printf("id: %s; page: %s; name: %s message: %s", id, page, name, message)
	})

	//上传文件
	//单文件
	router.POST("/upload", func(context *gin.Context) {
		//单文件
		file, handler, err := context.Request.FormFile("file")
		if err != nil {
			log.Fatalf("context.Request.FormFile err: %s", err)
		}
		fileName := handler.Filename
		log.Println(fileName)

		out, err := os.Create("./tmp/" + fileName)
		if err != nil {
			log.Fatalf("os.Create err: %s", err)
		}
		defer out.Close()
		//上传文件到指定的dst
		size, err := io.Copy(out, file)
		if err != nil {
			log.Fatalf("io.Copy err: %s", err)
		}

		context.String(http.StatusOK, fmt.Sprintf("%s' uploaded! size: %dbit", fileName, size))
	})

	//为 multipart 表单设置一个较低的内存限制（默认是32MiB）
	//8MiB
	//router.MaxMultipartMemory = 8 << 20
	//多文件上传
	router.POST("/upload2", func(context *gin.Context) {
		err := context.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatalf("context.Request.ParseMultipartForm err: %s", err)
		}
		//Multipart form
		form := context.Request.MultipartForm
		files := form.File["upload[]"]

		for index := range files {
			file, err := files[index].Open()
			defer file.Close()
			if err != nil {
				log.Fatalf("files[%d].Open err: %s", index, err)
			}
			log.Println(files[index].Filename)

			out, err := os.Create(files[index].Filename)
			defer out.Close()
			if err != nil {
				log.Fatalf("os.Create err: %s", err)
			}

			//上传文件到指定的dst
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatalf("io.Copy err: %s", err)
			}
		}

		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	//组路由
	//简单组：v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(context *gin.Context) {})
		v1.POST("/submit", func(context *gin.Context) {})
		v1.POST("/read", func(context *gin.Context) {})
	}

	//简单组：v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(context *gin.Context) {})
		v2.POST("/submit", func(context *gin.Context) {})
		v2.POST("/read", func(context *gin.Context) {})
	}

	//默认情况下，它使用：8080，除非定义了PORT环境变量
	err := router.Run()
	//硬编码端口
	//err := router.Run(":3000")
	if err != nil {
		log.Fatalf("router.Run err: %s", err)
	}

}
