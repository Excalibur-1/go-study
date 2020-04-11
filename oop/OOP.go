package main

import "fmt"

//模拟构造函数
//Go和传统的面向对象语言如Java有着很大区别。结构体没有构造函数初始化功能，可以通过以下方式模拟
type Person struct {
	Name string
	Age  int
}

func NewPersonByName(name string) *Person {
	return &Person{
		Name: name,
	}
}

//因为Go没有函数重载，为了避免函数名字冲突，使用了NewPersonByName和NewPersonByAge两个不同的函数表示不同的Person构造过程。
func NewPersonByAge(age int) *Person {
	return &Person{
		Age: age,
	}
}

//父子关系结构体初始化
type Student struct {
	Person
	ClassName string
}

//构造父类
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

//构造子类
func NewStudent(className string) *Student {
	p := &Student{}
	p.ClassName = className
	return p
}

//在go中，可以给任意类型（除了指针）添加方法
type Integer int

func (i Integer) Less(j Integer) bool {
	return i < j
}

//Golang 中的方法是作用在指定的数据类型上的(即:和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅仅是 struct。
//方法的声明和调用
/*func (receiver 类型) methodName(参数列表) (返回值列表) {
	//方法体
	return 返回值列表
}*/

//方法与函数的示例
//一个run函数
func run(p *Person, name string) {
	p.Name = name
	fmt.Println("函数 run...", p.Name)
}

//一个run方法
func (p *Person) run() {
	fmt.Println("方法 run...", p.Name)
}

/*
Go方法本质
Go的方法是一种作用于特定类型变量的函数，这种特定类型的变量叫做接收器（Receiver）。如果特定类型理解为结构体或者“类”时，接收器就类似于其他语言的this或者self。
在Go中，接收器可以是任何类型，不仅仅是结构体，依此我们看出，Go中的方法和其他语言的方法类似，但是Go语言的接收器强调方法的作用对象是实例。
方法与函数的区别就是：函数没有作用对象。
上述Person案例中，接收器类型是*Person，属于指针类型，非常接近Java中的this，由于指针的特性，调用方法时，修改接收器指针的任意长远变量，在方法结束后，修改都是有效的。
当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效，如下所示：
*/

//定义一个表示点的结构体
type Point struct {
	X int
	Y int
}

//非指针接收器
func (p Point) Add(otherP Point) Point {
	return Point{
		p.X + otherP.X,
		p.Y + otherP.Y,
	}
}

func main() {
	p := NewPersonByName("zs")
	fmt.Println(p)

	s := NewStudent("一班")
	fmt.Println(s)

	//调用自定义类型绑定方法
	var i Integer = 1
	fmt.Println(i.Less(5))

	//实例化一个对象
	p1 := &Person{
		Name: "ruyue",
		Age:  10,
	}

	//执行函数：输出：函数 run... 张三
	run(p1, "张三")

	//执行方法，输出：方法 run... 张三，因为函数run修改了p1的name
	p1.run()

	//非指针接收器测试
	p2 := Point{1, 1}
	p3 := Point{2, 2}
	//不会修改原来点的位置
	result := p2.Add(p3)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(result)
	//一般情况下，小对象由于复制时速度较快，适合使用非指针接收器，
	//大对象因为复制性能较低，适合使用指针接收器，
	//此时再接收器和参数之间传递时不进行复制， 只传递指针。
}
