package main

import (
	"fmt"
	"study/oop/person"
)

type Father struct {
	Name string
	age  int
}

func (f *Father) run() {
	fmt.Println(f.Name + " like running...")
}

//继承
type Son struct {
	Father //匿名嵌套结构体
}

/*
注意：
当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。
结构体嵌入多个匿名结构体，如果两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，访问时必须明确指定匿名结构体名字，否则编译报错。
如果一个 struct 嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字。
*/

//关于多重继承：如果一个 struct 嵌套了多个匿名结构体，那么该struct可以直接访问嵌套的匿名结构体的字段和方法，从而实现多重继承。
type Father2 struct {
	Like string
}

type Son1 struct {
	Father
	Father2
}

type Son2 struct {
	*Father
	*Father2
}

func main() {
	//调用person包下的方法，演示面向对象的封装特性
	p := person.NewPerson("Tom")
	p.SetAge(18)
	fmt.Println(p)

	var s Son
	/*s.Father.Name = "Tom"
	//可以访问未导出属性
	s.Father.age = 10
	//可以访问未导出方法
	s.Father.run()*/

	//上述可以简写为：
	s.Name = "Tom"
	s.age = 10
	s.run()

	s1 := &Son1{
		Father: Father{
			Name: "Tom",
			age:  10,
		},
		Father2: Father2{
			Like: "伏特加",
		},
	}
	fmt.Println(s1)

	s2 := &Son2{
		&Father{
			Name: "Tom",
			age:  10,
		},
		&Father2{
			Like: "伏特加",
		},
	}
	fmt.Println(s2.Father)

	//多态，参见接口测试包
}
