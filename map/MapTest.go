package main

import (
	"fmt"
	"sync"
)

type Person struct {
	ID   string
	Name string
}

func main() {
	//创建map
	//[]内的类型为任意可以进行比较的类型 int是值（value）的类型
	m1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(m1)
	//make方式创建map
	var m2 map[string]Person
	m2 = make(map[string]Person)
	//添加值
	m2["123"] = Person{"123", "Tom"}
	//判断是否存在key为"123"的键值，并返回key-value
	p, isFind := m2["123"]
	fmt.Println(isFind)
	fmt.Println(p)
	//注意：golang中map的 key 通常 key 为 int 、string，
	//但也可以是其他类型如：bool、数字、string、指针、channel，
	//还可以是只包含前面几个类型的接口、结构体、数组。
	//slice、map、function不能使用 == 来判断，不能作为map的key。

	//map使用
	//通过key操作元素
	var numbers map[string]int
	numbers = make(map[string]int)
	//赋值
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	//删除key为ten的元素
	delete(numbers, "ten")
	//读取数据
	fmt.Println("第三个数字是：", numbers["three"])

	//map遍历,与数组、slice一样使用for-range结构遍历
	/*
		map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取；
		map的长度是不固定的，也就是和slice一样，也是一种引用类型
		内置的len函数同样适用于map，返回map拥有的key的数量
		go没有提供清空元素的方法，可以重新make一个新的map，不用担心垃圾回收的效率，因为go中并行垃圾回收效率比写一个清空函数高效很多
		map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	*/

	//并发安全的map
	//下面的代码会报错，因为对map进行了并发读写，读写map是线程不安全的
	/*m := make(map[int]int)
	go func() {
		for {
			//无限写入
			m[1] = 1
		}
	}()

	go func() {
		for {
			//无限读取
			_ = m[1]
		}
	}()

	for{}*/

	//go1.9新增了sync.Map，高效且并发安全
	/*
		无须初始化，直接声明即可
		sync.Map不能使用map的方式进行取值和设值操作，而是使用sync.Map的方法进行调用。Store表示存储，Load表示获取，Delete表示删除。
		使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，需要继续迭代时，返回true，终止迭代返回false。
	*/
	var scene sync.Map

	//保存键值对
	scene.Store("id", 1)
	scene.Store("name", "lisi")

	//根据键取值
	fmt.Println(scene.Load("name"))

	//遍历
	scene.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
