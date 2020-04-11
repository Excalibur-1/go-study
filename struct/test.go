package main

import "fmt"

//定义结构体
type person struct {
	Name string
	Age  int
}

type person1 struct {
	Name string
	Age  int
	//结构内部定义匿名结构
	Contact struct {
		Phone, City string
	}
}

//匿名字段，字面值初始化必须按照声明顺序进行赋值
type person2 struct {
	string
	int
}

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
	Sex  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := person{}
	a.Name = "Joe"
	a.Age = 19
	fmt.Println(a)
	//字面值初始化写法
	b := person{Name: "joe", Age: 19}
	fmt.Println(b)

	//此处传递的依然是值的拷贝
	A(b)
	fmt.Println(b)

	//如果需要修改原对象的值则可以采用传递指针的形式
	B(&b)
	fmt.Println(b)

	//当需要重复对原结构体进行修改时，指针的获取可以直接赋值给一个变量
	c := &person{Name: "Joe", Age: 19}
	B(c)
	//修改结构体属性时，不需要使用"*"指定，可以直接修改，这与其他语言不同，开发过程中建议直接使用取地址方式初始化结构体
	c.Name = "OK"
	fmt.Println(c)

	//匿名结构
	d := &struct {
		Name string
		Age  int
	}{Name: "joe", Age: 19}
	fmt.Println(d)

	//结构体嵌套匿名结构
	e := &person1{Name: "Joe", Age: 19}
	e.Contact.Phone = "1231231241"
	e.Contact.City = "shenzhen"
	fmt.Println(e)

	//对匿名字段结构进行初始化
	f := person2{"Joe", 19}
	fmt.Println(f)

	a = person{Name: "Joe", Age: 19}
	b = person{Name: "Joe", Age: 19}
	//相同类型的比较，值相同比较的结果相同，类型不同不能进行比较
	fmt.Println(a == b)

	//匿名嵌入结构的字面值初始化方式
	g := teacher{Name: "Joe", Age: 19, human: human{Sex: 0}}
	h := student{Name: "Joe", Age: 19, human: human{Sex: 1}}
	g.Name = "Joe2"
	g.Age = 13
	//两种赋值方式，第二种方式可以用来解决匿名字段和外层结构存在同名字段的问题
	g.Sex = 100
	g.human.Sex = 100
	h.human.Sex = 200
	fmt.Println(g, h)
}

func A(per person) {
	per.Age = 20
	fmt.Println("A", per)
}

func B(per *person) {
	per.Age = 20
	fmt.Println("A", per)
}
