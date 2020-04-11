package main

import (
	"fmt"
	"sort"
)

func main() {
	//声明并创建map的几种方式
	var m map[int]string
	m = make(map[int]string)
	fmt.Println(m)

	var m1 map[int]string = make(map[int]string)
	fmt.Println(m1)

	var m2 = make(map[int]string)
	fmt.Println(m2)

	m3 := make(map[int]string)
	fmt.Println(m3)

	//向map存值
	m[1] = "OK"
	//取值，如果不存在则为空
	a := m[1]
	fmt.Println("map：", m, a)
	//删除键值
	delete(m, 1)
	b := m[1]
	fmt.Println("map：", m, b)

	//复杂map， 多层map嵌套时需要注意每次对内层map进行操作时都需要检查，才能避免异常
	var m4 map[int]map[int]string
	m4 = make(map[int]map[int]string)
	/*
		此处需要对内层map进行初始化操作，否则会提示nil异常，
		但是此处只是针对外层map的key=1的内层map进行初始化，
		如果调用map[2][1]仍然会抛出异常
	*/
	m4[1] = make(map[int]string)
	m4[1][1] = "OK"
	c := m4[1][1]
	fmt.Println(c)

	//使用多参数返回进行校验
	d, ok := m4[2][1]
	fmt.Println(d, ok)
	//如果不存在，进行初始化操作
	if !ok {
		m4[2] = make(map[int]string)
	}
	m4[2][1] = "GOOD"
	d, ok = m4[2][1]
	fmt.Println(d, ok)

	//迭代map
	//使用slice嵌套map
	sm := make([]map[int]string, 5)
	//此处循环对map进行赋值，不会影响slice，因为这里采用的是拷贝形式
	for _, v := range sm {
		v = make(map[int]string)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)

	//此处采用slice下标的形式赋值，可以实现对slice中的map进行赋值操作
	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	//对map进行间接排序
	map1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "d"}
	slice1 := make([]int, len(map1))
	//将map中所有的key存入slice
	i := 0
	for k := range map1 {
		slice1[i] = k
		i++
	}
	fmt.Println(slice1)
	//调用sort函数对slice排序
	sort.Ints(slice1)
	//排序后即可根据slice元素的顺序取出map的值
	fmt.Println(slice1)

	//交换map的键值
	changeMap()
}

//交换map的键值
func changeMap() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
