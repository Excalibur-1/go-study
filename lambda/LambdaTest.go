package main

import "fmt"

/*
闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使己经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量。
简单的说 : 函数+引用环境=闭包
贴士：闭包( Closure)在某些编程语言中也被称为 Lambda表达式（如Java）
*/
func main() {
	//在闭包中可以修改引用的变量
	str := "hello"
	foo := func() {
		str = "world"
	}
	foo()
	//输出：world
	fmt.Println(str)

	f := fn1(1)
	//输出地址
	fmt.Println(f)
	g := fn1(2)
	//输出地址
	fmt.Println(g)

	//输出:地址,1,1
	fmt.Println(",", f(1))
	//输出:地址,2,2
	fmt.Println(",", g(1))

	//输出:地址,1,1
	fmt.Println(",", f(2))
	//输出:地址,2,2
	fmt.Println(",", g(2))

	//案例2，累加器
	//返回一个闭包
	accAdd := Accumulate(1)
	//输出：2
	fmt.Println(accAdd())
	//输出：3
	fmt.Println(accAdd())

}

//简单示例
func fn1(a int) func(i int) int {
	return func(i int) int {
		print(&a, ",", a)
		return a
	}
}

//累加器
func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
