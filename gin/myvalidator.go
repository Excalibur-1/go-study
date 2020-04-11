package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binging:"required,bookableddate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binging:"required,gtfield=CheckIn" time_format:"2006-01"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func getBookable(context *gin.Context) {
	var b Booking
	if err := context.ShouldBindWith(&b, binding.JSON); err == nil {
		context.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("bookabledate", bookableDate)
		if err != nil {
			log.Fatalf("v.RegisterValidation err: %s", err)
		}
	}

	router.GET("/bookable", getBookable)

	//默认情况下，它使用：8080，除非定义了PORT环境变量
	err := router.Run()
	//硬编码端口
	//err := router.Run(":3000")
	if err != nil {
		log.Fatalf("router.Run err: %s", err)
	}

}
