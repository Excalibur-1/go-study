package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c) // 第一行 - c 的值为 31
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c) // 第二行 - c 的值为 11
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c) // 第三行 - c 的值为 210
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c) // 第四行 - c 的值为 2
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c) // 第五行 - c 的值为 1
	a++
	fmt.Printf("第六行 - c 的值为 %d\n", a) // 第六行 - c 的值为 22
	a--
	fmt.Printf("第七行 - c 的值为 %d\n", a) // 第七行 - c 的值为 21
}
