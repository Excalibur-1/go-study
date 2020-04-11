//反射是指在程序运行期对程序本身进行访问和修改的能力，即可以在运行时动态获取变量的各种信息，比如变量的类型（type），类别（kind），如果是结构体变量，还可以获取到结构体本身的信息（字段与方法），通过反射，还可以修改变量的值，可以调用关联的方法。
//反射常用在框架的开发上，一些常见的案例，如JSON序列化时候tag标签的产生，适配器函数的制作等，都需要用到反射。反射的两个使用常见使用场景：
//不知道函数的参数类型：没有约定好参数、传入类型很多，此时类型不能统一表示，需要反射
//不知道调用哪个函数：比如根据用户的输入来决定调用特定函数，此时需要依据函数、函数参数进行反射，在运行期间动态执行函数
//Go程序的反射系统无法获取到一个可执行文件空间中或者是一个包中的所有类型信息，需要配合使用标准库中对应的词法、语法解析器和抽象语法树( AST) 对源码进行扫描后获得这些信息。
//贴士：
//C，C++没有支持反射功能，只能通过 typeid 提供非常弱化的程序运行时类型信息。
//Java、 C#等语言都支持完整的反射功能。
//Lua、JavaScript类动态语言，由于其本身的语法特性就可以让代码在运行期访问程序自身的值和类型信息，因此不需要反射系统。
//注意：
//在编译期间，无法对反射代码进行一些错误提示。
//反射影响性能
package main

import (
	"fmt"
	"reflect"
)

//反射是通过接口的类型信息实现的，即反射建立在类型的基础上：当向接口变量赋予一个实体类型的时候，接口会存储实体的类型信息。
//Go中反射相关的包是reflect，在该包中，定义了各种类型，实现了反射的各种函数，通过它们可以在运行时检测类型的信息、改变类型的值。
//变量包括type、value两个部分（所以 nil != nil ），type包括两部分：
//static type：在开发时使用的类型，如int、string
//concrete type：是runtime系统使用的类型
//类型能够断言成功，取决于 concrete type ，如果一个reader变量，如果 concrete type 实现了 write 方法，那么它可以被类型断言为writer。
//Go中，反射与interface类型相关，其type是 concrete type，只有interface才有反射！每个interface变量都有一个对应的pair，pair中记录了变量的实际值和类型（value, type）。即一个接口类型变量包含2个指针，一个指向对应的 concrete type ，另一个指向实际的值 value。

//反射初识
//reflect包的2个函数：
//ValueOf()：获取变量的值，即pair中的 value
//TypeOf()：获取变量的类型，即pair中的 concrete type
type Person struct {
	Name string
	Age  int
}

type user struct {
	Name string
	Age  int `json:"age" id:"100"` //结构体标签
}

func (u *user) AddAge(addNum int) {
	fmt.Println("age add result:", u.Age+addNum)
}

func (u *user) ShowName() {
	fmt.Println(u.Name)
}

func main() {
	//var r io.Reader                                   //定义了一个接口类型
	//r, err := os.OpenFile("", os.O_RDWR, os.ModePerm) //记录接口类型的实际类型、实际值

	//var w io.Writer   //定义一个接口类型
	//w = r.(io.Writer) //赋值时，接口内部的pair不变，所以 w 和 r 是同一类型

	p := Person{"lisi", 31}

	fmt.Println(reflect.ValueOf(p)) //变量的值

	fmt.Println(reflect.ValueOf(p).Type()) //变量类型的对象名
	fmt.Println(reflect.TypeOf(p))         //变量类型的对象名)

	fmt.Println(reflect.TypeOf(p).Name()) //变量类型对象的类型名
	fmt.Println(reflect.TypeOf(p).Kind()) //变量类型对象的种类名

	fmt.Println(reflect.TypeOf(p).Name() == "Person")       //true
	fmt.Println(reflect.TypeOf(p).Kind() == reflect.Struct) //true

	//类型与种类的区别
	/*
		Type是原生数据类型： int、string、bool、float32 ，以及 type 定义的类型，对应的反射获取方法是 reflect.Type 中 的 Name()
		Kind是对象归属的品种：Int、Bool、Float32、Chan、String、Struct、Ptr（指针）、Map、Interface、Func、Array、Slice、Unsafe Pointer等
	*/

	//静态类型与动态类型
	//静态类型：变量声明时候赋予的类型
	type MyInt int //int是静态类型
	var i *int     //*int是静态类型

	var A interface{} //空接口是静态类型，必须是接口类型才能实现类型动态变化
	A = 10            //此时静态类型为interface{} 动态类型为int
	A = "hello"       //此时静态类型为interface{} 动态类型为string
	fmt.Println(&i, A)

	//反射使用
	//反射操作简单数据类型
	var num int64 = 100
	//设置值：指针传递
	ptrValue := reflect.ValueOf(&num)
	newValue := ptrValue.Elem()                //Elem()用于获取原始值的反射对象
	fmt.Println("type:", newValue.Type())      //int64
	fmt.Println("can set:", newValue.CanSet()) //true
	newValue.SetInt(200)
	fmt.Println(newValue)

	//获取值：值传递
	rValue := reflect.ValueOf(num)
	fmt.Println(rValue.Int())               //方式一：200
	fmt.Println(rValue.Interface().(int64)) //方式二：200

	//反射进行类型推断
	u := &user{
		Name: "Ruyue",
		Age:  100,
	}

	fmt.Println(reflect.TypeOf(u))         //*main.user
	fmt.Println(reflect.TypeOf(*u))        //main.user
	fmt.Println(reflect.TypeOf(*u).Name()) //user
	fmt.Println(reflect.TypeOf(*u).Kind()) //struct

	//反射操作指针
	typeOfUser := reflect.TypeOf(u).Elem()
	fmt.Println("element name:", typeOfUser.Name()) //user
	fmt.Println("element kind:", typeOfUser.Kind()) //struct

	//反射操作结构体，反射可以获取结构体的详细信息：
	//字段用法
	for i := 0; i < typeOfUser.NumField(); i++ { // NumField 当前结构体有多少个字段
		fieldType := typeOfUser.Field(i) // 获取每个字段
		fmt.Println(fieldType.Name, fieldType.Tag)
	}

	if userAge, ok := typeOfUser.FieldByName("Age"); ok {
		fmt.Println(userAge)
	}

	//方法用法
	for i := 0; i < typeOfUser.NumMethod(); i++ {
		fieldType := typeOfUser.Method(i) //获取每个字段
		fmt.Println(fieldType.Name)
	}

	//反射调用函数与方法
	/*
		 使用反射调用函数
		如果反射值对象(reflect.Value)中值的类型为函数时，可以通过 reflect.Value调用该 函数。
		使用反射调用函数时，需要将参数使用反射值对象的切片 reflect.Value 构造后传入 Call()方法中 ，
		调用完成时，函数的返回值通过 []reflect.Value 返回 。
	*/
	funcValue := reflect.ValueOf(add)
	params := []reflect.Value{reflect.ValueOf("lisi"), reflect.ValueOf(20)}
	reList := funcValue.Call(params)
	fmt.Println(reList) //函数返回值

	//反射调用方法
	//方法的调用是需要接受者的
	v := reflect.ValueOf(u)

	//调用无参方法
	methodV := v.MethodByName("ShowName")
	//methodV.Call(nil)
	//或者传递空切片也可以
	var value []reflect.Value
	methodV.Call(value)

	//调用有参方法
	methodV2 := v.MethodByName("AddAge")
	args := []reflect.Value{reflect.ValueOf(30)}
	methodV2.Call(args)
}

func add(name string, age int) string {
	fmt.Printf("name is %s, age is %d\n", name, age)
	return "success"
}
