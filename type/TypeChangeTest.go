package main

import (
	"fmt"
	"strconv"
	"study/type/mypack"
)

//使用不同包下的类型
type Student mypack.Person

//如果将类型定义改为类型别名：
//这时Student的方法就会报错：无法为 Person 添加新的方法
//type Student = mypack.Person

func (s *Student) Study() {
	fmt.Println("study...")
}

//命名类型，其类型为Person
type Person struct {
	name string
}

func main() {
	//go在不同类型变量之间需要显示转换
	//数值类型转换
	var i int = 100
	var f float64 = float64(i)
	fmt.Printf("f = %f\n", f)

	//基本数据类型与字符串转换
	var b bool = true
	var str string

	//fmt.Sprintf()根据format参数返回转换后的字符串
	str = fmt.Sprintf("%t", b)
	fmt.Println(str)

	//字符串转基本数据类型，如果转换的数据不合法则go默认会置为对应类型的零值
	var str1 string = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("%t\n", b1)

	//类型别名
	//go1.9 新增类型 type byte uint8 和 type rune int32
	//go1.9 之后使用类型别名 type byte = uint8 和 type rune = int32

	//类型定义是定义了一个全新的类型，类型别名只是某个类型的小名并非创造了新的类型
	//类型定义
	type MyInt int
	//类型别名
	type AliasInt = int

	var a1 MyInt
	fmt.Printf("a1 type:%T\n", a1)
	var a2 AliasInt
	fmt.Printf("a2 type:%T\n", a2)

	s := &Student{}
	s.Study()
	person := &mypack.Person{}
	person.Run()
	//调用类型别名的方法
	person.Study()

	//命名类型，见Person类型

	//未命名类型
	p := struct {
		name string
	}{}
	fmt.Println(p)

	//底层类型
	//预声明类型（Pre-declared types）和类型字面量（type literals）的底层类型是他们自身
	//自定义类型type newtype oldtype中newtype的底层类型是逐层递归向下查找的，直到找到oldtype的预声明类型或字面量类型

	//go中的类型相同
	/*
		命名类型的数据类型相同：声明语句必须完全相同
		未命名类型数据类型相同：类型声明字面量结构相同，且内部元素的类型相同
		命名类型与未命名类型永远不同
		通过类型别名语句声明的两个类型相同。类型别名语法：type T1 = T2
	*/

}
