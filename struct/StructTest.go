package main

import "fmt"

type Course []string

type Person struct {
	Name string
	Age  int
}

type Student struct {
	//匿名字段，Student默认包含了Person的所有字段
	Person
	//内置切片类型
	Course
	ClassName string
}

func main() {
	//顺序初始化，每个成员都必须初始化
	var p1 Person = Person{"lisi", 20}
	fmt.Println(p1)

	//指定成员初始化，没有被初始化的自动赋零值
	p2 := Person{Age: 30}
	fmt.Println(p2)

	//new 申请结构体
	//被new生成的结构体实例是指针类型
	p3 := new(Person)
	//这里的.语法只是GO语法糖，将p3.name转换成了(*p3).name
	p3.Name = "zs"
	p3.Age = 27
	fmt.Println(p3)

	//直接声明
	var p4 Person
	p4.Name = "ww"
	p4.Age = 30
	fmt.Println(p4)

	//结构体地址与实例化

	//使用&操作符进行取地址操作可以看做一次new实例化操作
	p5 := &Person{}
	fmt.Println(p5)

	//内嵌结构体
	//初始化方式1
	s1 := Student{
		Person{Age: 15, Name: "xm"},
		[]string{"80", "70"},
		"一班",
	}
	fmt.Println(s1.Age)
	fmt.Println(s1.Person.Age)

	//初始化方式2
	var s2 Student
	s2.Name = "xh"
	s2.Age = 14
	s2.ClassName = "二班"
	fmt.Println(s2.Age)
	fmt.Println(s2.Age)

	//匿名字段，见struct student

	//创建一个学生
	s3 := Student{Person{"xq", 16}, []string{"90"}, "三班"}

	//访问该学生的字段
	fmt.Println("name=", s3.Name)
	fmt.Println("className=", s3.ClassName)

	//修改学生的年龄
	s3.Age = 17
	fmt.Println("Age=", s3.Age)

}
