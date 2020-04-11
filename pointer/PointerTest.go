package main

import "fmt"

func main() {
	//Go保留了指针，代表某个内存地址，默认值为 nil ，使用 & 取变量地址，通过 * 访问目标对象。
	var a int = 10
	//一个16进制数
	fmt.Println("&a=", &a)

	var p *int = &a
	//10
	fmt.Println("*p=", *p)

	/*
		Go同样支持多级指针，如 **T
		空指针：声明但未初始化的指针
		野指针：引用了无效地址的指针，如：var p *int = 0，var p *int = 0xff00(超出范围)
		Go中直接使用.访问目标成员
	*/

	//指针使用示例
	b := 5
	swap(&a, &b)
	fmt.Println(a, b)

	//结构体指针
	var u = User{
		Name: "lisi",
		Age:  18,
	}
	p1 := &u
	//输出一样
	fmt.Println(u.Name)
	fmt.Println(p1.Name)

	//go不支持指针运算
	/*a1 := 1
	p2 := &a1
	//报错non-numeric type *int
	p2 ++*/

	//new()函数使用
	var p2 *bool
	//new()函数可以在 heap堆区申请一片内存地址空间：
	p2 = new(bool)
	//false
	fmt.Println(*p2)
}

type User struct {
	Name string
	Age  int
}

//指针使用示例：实现变量值交换
func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}
