package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

//模拟一些私有的数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

//自定义模板函数
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//设置简单的变量
		c.Set("example", "12345")

		//在请求之前
		log.Print(t)

		c.Next()

		//在请求之后
		latency := time.Since(t)
		log.Print(latency)

		//记录我们的访问状态
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.Default()

	// gin.H 是一个 map[string]interface{} 的快捷方式
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		//你也可以使用一个结构
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		//注意 msg.Name 在JSON中会变成"user"
		//将会输出： {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//SecureJSON
	//使用 SecureJSON 来防止 json 劫持。如果给定的结构体是数组值，默认预置 "while(1)," 到 response body 。
	//你也可以使用自己安装的json前缀
	//r.SecureJsonPrefix(")]}',\n")
	r.GET("/someSecureJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		//将会输出：while(1);["lena","austin","foo"]
		//SecureJSON在新版本貌似没有了。。。
		//c.SecureJSON(http.StatusOK, names)
		c.JSON(http.StatusOK, names)
	})

	//JSONP
	//在不同的域中使用 JSONP 从一个服务器请求数据。
	//如果请求参数中存在 callback，添加 callback 到 response body 。
	r.GET("JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback 是 x
		//将会输出： x({\"foo\":\"bar\"})，此版本貌似没有这个api
		//c.JSONP(http.StatusOK, data)
		c.JSON(http.StatusOK, data)
	})

	//AsciiJSON
	//使用 AsciiJSON 生成仅有 ASCII 字符的 JSON，非 ASCII 字符将会被转义 。
	r.GET("AsciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 将会输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		//c.AsciiJSON(http.StatusOK, data)
		c.JSON(http.StatusOK, data)
	})

	//静态文件服务
	r.Static("/asserts", "./asserts")
	r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	//从reader提供数据
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		//reader := response.Body
		//contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		/*extraHeaders := map[string]string{
			"Content-Disposition":`attachment;filenamegopher.png=`,
		}*/

		//c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
		c.Data(http.StatusOK, contentType, []byte(""))
	})

	//HTML 渲染
	//使用 LoadHTMLGlob () 或 LoadHTMLFiles ()
	//r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	/*r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":"Main website",
		})
	})*/

	//在不同的目录使用具有相同名称的模板
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	//自定义模板渲染器
	//html := template.Must(template.ParseFiles("file1", "file2"))
	//r.SetHTMLTemplate(html)

	//自定义分隔符
	r.Delims("{[{", "}]}")
	//r.LoadHTMLGlob("/path/to/templates")

	//使用自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("./fixtures/basic/raw.tmpl")

	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	//重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	//路由重定向
	r.GET("test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	//使用自定义中间件
	r.Use(Logger())
	r.GET("/test3", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		//它将打印：“12345”
		log.Println(example)
	})

	//在组中使用gin.BasicAuth() 中间件
	//gin.Accounts 是map[string]string 的快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 结尾
	//点击 “localhost:8080/admin/secrets”
	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET:("})
		}
	})

	//在中间件中使用协程
	//在一个中间件或处理器中启动一个新的协程时，你 不应该 使用它里面的原始的 context ，只能去使用它的只读副本。
	r.GET("long_async", func(c *gin.Context) {
		//创建在协程中使用的副本
		cCp := c.Copy()
		go func() {
			//使用 time.Sleep() 休眠5秒，模拟一个用时长的任务
			time.Sleep(5 * time.Second)

			//注意，你使用的是复制的 context “cCp”，重要
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 使用 time.Sleep() 休眠 5 秒，模拟一个用时长的任务。
		time.Sleep(5 * time.Second)

		// 因为我们没有使用协程，我们不需要复制 context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	//自定义HTTP配置
	//直接使用http.ListenAndServe()，像这样
	/*err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(" http.ListenAndServe err:", err)
	}*/

	//或
	/*s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(" http.ListenAndServe err:", err)
	}*/

	//支持 Let's Encrypt
	//一个 LetsEncrypt HTTPS 服务器的示例。
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	/*log.Fatal(autotls.Run(r, "example1.com", "example2.com"))*/

	//自定义 autocert 管理器示例。
	/*m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))*/

	err := r.Run()
	if err != nil {
		log.Fatal("r.Run err:", err)
	}
}
