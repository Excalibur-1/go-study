package main

import "fmt"

func Add01(a, b int) int {
	return a + b
}

type long int

func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	var result1 int
	//普通函数调用方法
	result1 = Add01(1, 1)
	fmt.Println("result1 = ", result1)

	var a long = 2
	//面向对象只是换了种表现形式
	r := a.Add02(2)
	fmt.Println("r = ", r)

}
