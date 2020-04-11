package main

import "fmt"

func main() {
	// 这里不能写成 b := []byte{"Golang"}，这里是利用类型转换。
	b := []byte("Golang")
	c := []byte("go")
	d := []byte("Go")
	fmt.Println(b, c, d)
}
