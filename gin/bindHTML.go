package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func startPage(c *gin.Context) {
	var person Person

	//如果是 GET，只使用 Form 绑定引擎 query
	//如果是 POST，首先检查 content-type 为 JSON 或 XML 然后使用 Form（form-data）
	if c.ShouldBindWith(&person, binding.Form) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else if c.ShouldBindWith(&person, binding.JSON) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else if c.ShouldBindWith(&person, binding.XML) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else if c.ShouldBindWith(&person, binding.FormPost) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	/*if c.ShouldBindWith(&person, binding.JSON) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}*/
	c.String(http.StatusOK, "Success")
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	if c.ShouldBindWith(&fakeForm, binding.Form) == nil {
		log.Println(fakeForm.Colors)
	}
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

func main() {
	route := gin.Default()

	route.Any("/testing", startPage)
	route.POST("/", formHandler)

	err := route.Run()
	if err != nil {
		log.Fatalln("route.Run err: ", err)
	}
}
