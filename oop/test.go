package main

import "fmt"

type A struct {
	Name string
	//首字母小写表示访问权限仅在本包内，其他包访问不到此属性
	name1 string
}

type B struct {
	Name string
}

//别名类型定义
type TZ int

type ME int

type ME2 struct {
	a int
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	b := B{}
	b.Print()
	fmt.Println(b.Name)

	//底层类型的方法绑定
	var c TZ
	c.Print()
	//另一种方法调用方式
	(*TZ).Print(&c)

	var d ME
	d.Increase()
	fmt.Println(d)
	d.Increase()
	fmt.Println(d)
	d.Increase()
	fmt.Println(d)

	e := ME2{}
	e.Increase()
	fmt.Println(e)
	e.Increase()
	fmt.Println(e)
	e.Increase()
	fmt.Println(e)
}

//定义方法
func (a *A) Print() {
	//此处利用指针可以修改属性的值
	a.Name = "AA"
	fmt.Println("A")
}

//go中没有方法重载概念，但是类型是不同的，即方法名虽然一样，但是实际上是不同的方法
func (b B) Print() {
	//此处针对b的属性进行的修改对外层调用不生效，因为修改的是拷贝的值
	b.Name = "BB"
	fmt.Println("B")
}

//对底层类型附带方法，同包中才可以绑定
func (a *TZ) Print() {
	fmt.Println("TZ")
}

//自增100方法
func (a *ME) Increase() {
	*a = *a + 100
}

func (a *ME2) Increase() {
	a.a += 100
}
