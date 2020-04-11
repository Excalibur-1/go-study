package main

import "fmt"

type myStr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型没有名字，匿名字段，继承了Person的成员
	int    //基础类型的匿名字段
	myStr  //自定义类型匿名字段
}

func main() {
	s := Student{Person{"mike", 'm', 16}, 1, "呵呵呵"}
	fmt.Printf("s = %+v\n", s)

	//成员赋值
	s.Person = Person{"go", 'm', 16}

	//非结构体匿名字段的打印
	fmt.Println(s.Person, s.int, s.myStr)
	fmt.Println(s.name, s.age, s.sex, s.int, s.myStr)
}
