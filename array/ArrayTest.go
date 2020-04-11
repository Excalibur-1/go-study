package main

import (
	"fmt"
	"strconv"
)

func main() {
	//初始化形式
	var arr1 [10]int
	//定义并初始化
	arr2 := [5]int{1, 2, 3, 4, 5}
	//未初始化的自动使用零值
	arr3 := [5]int{1, 2}
	//指定下标初始化
	arr4 := [5]int{2: 10, 4: 11}
	//自动计算长度
	arr5 := [...]int{2, 3, 4}
	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//常用操作
	//代表所有元素
	fmt.Println(arr1[:])
	//代表前五个元素，左闭右开
	fmt.Println(arr1[:5])
	//代表从第6个开始取到最后，左开右闭
	fmt.Println(arr1[5:])
	//数组长度
	fmt.Println(len(arr1))

	//数组遍历
	var formatInt string
	for i := 0; i < len(arr1); i++ {
		formatInt += strconv.FormatInt(int64(arr1[i]), 10) + " "
	}
	fmt.Println(formatInt)

	for k, v := range arr1 {
		fmt.Println(k, v)
	}

}
