package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (tmp *Person) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

//集成了Person的成员和方法
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

func (tmp *Student) PrintInfo() {
	fmt.Println("Student :tmp = ", tmp)
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	//就近原则，先找本作用域的方法，找不到再用继承的方法
	s.PrintInfo()

	//显示调用继承的方法
	s.Person.PrintInfo()
}
