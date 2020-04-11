package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
	name   string //和Person的那么同名
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 1, "bj", ""}

	//成员赋值
	s.name = "yoyo"
	s.sex = 'f'
	s.id = 666
	s.age = 22
	s.addr = "sz"

	//整体赋值
	s.Person = Person{"go", 'm', 16}

	//成员打印
	fmt.Println(s.name, s.sex, s.addr, s.age, s.id)

	//声明
	var a Student
	/*
		默认规则（就近原则）如果能在本作用域能找到此成员，就操作此成员
		如果没有找到，向上查找继承的字段，找到继承的字段
	*/
	a.name = "yoyo" //操作的是Student的name还是Person的name？
	a.sex = 'f'
	a.id = 666
	a.age = 22
	a.addr = "sz"
	fmt.Printf("a = %+v\n", a)

	//显示调用
	a.Person.name = "go" //Person的name
	fmt.Printf("a = %+v\n", a)
}
