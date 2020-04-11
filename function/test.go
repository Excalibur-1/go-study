package main

import "fmt"

func main() {
	fmt.Println(B())
	fmt.Println(C())
	a := 2
	//此处传递的是值的拷贝
	E(a)
	fmt.Println(a)
	I(1, 2)

	//将函数赋值给变量b
	b := K
	b()

	//匿名函数
	c := func() {
		fmt.Println("Func K")
	}
	c()

	//调用闭包函数
	d := closure(10)
	fmt.Println(d(1))
	fmt.Println(d(2))

	//defer函数，后进先出，即使发生严重错误也会执行，常用于资源清理、文件关闭等操作
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	//此处使用闭包，打印的值一直是循环的最后的一个值，因为使用的地址的引用
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	//Go中没有异常机制，采用panic/recover机制处理错误
	M()
	//此处执行了panic后程序会中止运行，通过使用defer + recover，将程序从panic状态恢复，程序才能继续执行
	N()
	P()

	test()

}

//函数定义，单个返回值参数
func A() int {
	a := 1
	return a
}

//多参数返回不指定返回参数名
func B() (int, int, int) {
	a, b, c := 1, 2, 3
	return a, b, c
}

//多参数返回指定返回参数名
func C() (a int, b int, c int) {
	//此处赋值不能使用":="因为变量已经在返回值中定义
	a, b, c = 1, 2, 3
	return
}

//多参数返回指定返回参数名，简写形式
func D() (a, b, c int) {
	//此处赋值不能使用":="因为变量已经在返回值中定义
	a, b, c = 1, 2, 3
	return
}

//单个参数函数定义
func E(a int) {
	a = 1
	fmt.Println(a)
}

//多个参数函数定义
func F(a int, b int) {
	a = 1
	fmt.Println(a)
}

//多个参数函数定义简写形式
func G(a, b int) {
	a = 1
	fmt.Println(a, b)
}

//多个参数函数定义，不同类型
func H(a int, b string) {
	a = 1
	fmt.Println(a, b)
}

//可变参数传递，实际上传递的是一个slice形式的值，传递的值是一个值的拷贝，必须放在参数列表的最后
func I(a ...int) {
	fmt.Println(a)
}

//次数直接使用slice作为参数传递，传递的是slice的地址
func J(a []int) {
	fmt.Println(a)
}

//函数作为参数赋值
func K() {
	fmt.Println("Func K")
}

//闭包，返回值是一个函数，这里定义的是一个参数为int返回值为int的函数
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func M() {
	fmt.Println("Func M")
}

func N() {
	//defer函数需要定义在panic之前
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in N")
}

func P() {
	fmt.Println("Func P")
}

func test() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer closure i =", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}
	for _, f := range fs {
		f()
	}
}
