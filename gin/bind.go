package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	//绑定JSON的示例 ({"user": "menu", "password": "123"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindWith(&json, binding.JSON); err == nil {
			if json.User == "menu" && json.Password == "123" {
				context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	//一个HTML表单绑定的示例(user=menu&password=123)
	router.POST("loginForm", func(context *gin.Context) {
		var form Login
		if err := context.ShouldBindWith(&form, binding.Form); err == nil {
			if form.User == "menu" && form.Password == "123" {
				context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	//Multipart/Urlencoded 绑定
	router.POST("/login", func(c *gin.Context) {
		var form Login
		// 你可以使用显示绑定声明绑定 multipart 表单：
		//c.ShouldBindWith(&form, binding.Form)
		// 或者你可以使用 ShouldBind 方法去简单的使用自动绑定(貌似新版本已经废弃)
		if c.ShouldBindWith(&form, binding.Form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(http.StatusOK, "you are logged in")
			} else {
				c.JSON(http.StatusUnauthorized, "unauthorized")
			}
		}
	})

	//默认情况下，它使用：8080，除非定义了PORT环境变量
	err := router.Run()
	//硬编码端口
	//err := router.Run(":3000")
	if err != nil {
		log.Fatalf("router.Run err: %s", err)
	}

}
