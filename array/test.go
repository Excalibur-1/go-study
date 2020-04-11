package main

import (
	"fmt"
)

func main() {
	//数组之间赋值必须大小一致
	var a [2]int
	var b [2]int
	b = a
	fmt.Println(b)

	//未赋值元素自动填充默认值
	c := [2]int{1}
	fmt.Println(c)

	//指定下标赋值
	d := [20]int{19: 1}
	fmt.Println(d)

	//不指定数组长度，自动根据元素数量计算长度
	e := [...]int{1, 2, 3, 4, 5}
	fmt.Println(e)

	//不指定长度的情况下，根据下标赋值，自动计算数组大小
	f := [...]int{19: 1}
	f[0] = 1
	fmt.Println(f)

	//指向数组的指针，指向数组f的地址
	var p *[20]int = &f
	fmt.Println(p)
	//另一种指向数组的指针创建方式
	q := new([10]int)
	//指向数组的指针也可以采用”数组名[下标] = 值“的赋值形式
	q[0] = 1
	fmt.Println(q)

	//指针数组，数组存储的是地址
	x, y := 1, 2
	g := [...]*int{&x, &y}
	fmt.Println(g)

	//注意：数组在go中是值类型

	//数组比较采用==和！=，必须同类型数组才行，区分类型不仅包含类型，也包含长度
	h := [2]int{1, 2}
	i := [2]int{1, 2}
	fmt.Println(h == i)

	//多维数组
	j := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(j)
	k := [...][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(k)

	bubble()
}

//冒泡排序
func bubble() {
	a := [...]int{5, 2, 6, 1, 7, 3, 9}
	fmt.Println(a)
	//此处将长度获取方式提出，提高执行效率，这样每次循环时就不会重复获取数组长度
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
