package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

type MyConstruct1 struct {
	title string
}

type MyConstruct struct {
	Id   int
	Name string
	Age  int
	MyConstruct1
}

func (c MyConstruct) MyMethod(name, title string) {
	fmt.Printf("MyMethod executor：name=%s title=%s\n", name, title)
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func (u User) Hello1(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}

//使用反射获取信息
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("type error")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	//反射获取字段信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	//反射获取方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("type error")
		return
	} else {
		v = v.Elem()
	}

	//判断是否有 Name 这个字段
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
		return
	}

	//判断是否为string类型
	if f.Kind() == reflect.String {
		f.SetString("ByeBye")
	}
}

func SetMyConstruct(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("type error")
		return
	} else {
		v = v.Elem()
	}

	//判断是否有 Name 这个字段
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
		return
	}

	//判断是否有MyConstruct1
	t := v.FieldByName("MyConstruct1")
	if !t.IsValid() {
		fmt.Println("Bad MyConstruct1")
		return
	}

	//判断是否有Age
	a := v.FieldByName("Age")
	if !a.IsValid() {
		fmt.Println("Bad Age")
		return
	}

	//判断是否为int类型
	if a.Kind() == reflect.Int {
		a.SetInt(20)
	}

	//判断是否为string类型
	if f.Kind() == reflect.String {
		f.SetString("Bob")
	}

	//判断是否为Struct
	//if t.Kind() == reflect.Struct {
	//	//t1 := t.FieldByName("MyConstruct1")
	//	t1 := t.FieldByIndex([]int{0, 0})
	//	if !t1.IsValid() {
	//		fmt.Println("Bad Title")
	//		return
	//	}
	//
	//	//判断是否为string类型
	//	if t1.Kind() == reflect.String {
	//		t1.SetString("MyTitle")
	//	}
	//}
}

func main() {
	u := User{1, "OK", 12}
	Info(u)

	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	//查看字段属性，判断是否为匿名字段
	fmt.Printf("%#v\n", t.Field(0))

	//取出匿名字段中的属性
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

	//通过传递指针进行反射修改
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x) //输出999

	u1 := User{1, "OK", 12}
	Set(&u1)
	fmt.Println(u1)

	u.Hello1("Joe")

	//反射调用方法
	v1 := reflect.ValueOf(u)
	mv := v1.MethodByName("Hello1")
	//传入参数
	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)

	my := MyConstruct{Id: 1, Name: "OK", Age: 13, MyConstruct1: MyConstruct1{"test"}}
	SetMyConstruct(&my)
	fmt.Println(my)

	v2 := reflect.ValueOf(my)
	my1 := v2.MethodByName("MyMethod")
	args1 := []reflect.Value{reflect.ValueOf("Wang"), reflect.ValueOf("MyTitle")}
	my1.Call(args1)
	fmt.Println(&my1)
}
