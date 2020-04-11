package main

import (
	"fmt"
	"time"
)

//对slice的修改操作建议使用返回值返回修改后的值，避免地址变化导致的问题
func Pingpang(s []int) []int {
	//此处的操作会修改slice的地址
	s = append(s, 3)
	return s
}

func main() {
	//创建slice，初始大小为0
	s := make([]int, 0)
	fmt.Println(s)
	s = Pingpang(s)
	fmt.Println(s)

	t := time.Now()
	//此处针对时间格式的转换一定要使用go提供的常量或者与其相同的值，否则时间会产生问题
	fmt.Println(t.Format(time.ANSIC))

	s1 := []string{"a", "b", "c"}
	for _, v := range s1 {
		//此处必须将v当做参数传递进去，否则因为for range+闭包导致传递的值是v的地址，所以不传递参数时会直接打印三次c
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	select {}
}
