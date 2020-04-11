package main

import (
	"fmt"
)

func main() {
	//声明slice
	var s1 []int
	fmt.Println(s1)

	a := [10]int{0: 1, 9: 1}
	fmt.Println(a)
	//取出后5个元素赋值给slice，用10是因为这种语法包含头不包含尾
	s2 := a[5:10]
	fmt.Println(s2)
	//从索引为5的元素开始取到最后
	s3 := a[5:]
	fmt.Println(s3)
	//取出前5个元素
	s4 := a[:5]
	fmt.Println(s4)

	/*
		使用make函数声明slice，参数1为元素类型，参数2为元素个数，参数3为初始化大小，
		当元素格式超过初始化大小时，go会重新分配内存，分配大小为原来的2倍
		不指定第三个参数时，默认大小为元素个数
	*/
	s5 := make([]int, 3, 10)
	//len获取slice元素个数，cap获取slice容量
	fmt.Println(len(s5), cap(s5))

	//ReSlice
	b := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'i', 'j'}
	sa := b[2:5]
	fmt.Println(string(sa), len(sa), cap(sa))
	//从slice中取slice，下标以目标slice下标计算
	sb := sa[1:3]
	/*
		如果取以数组下标取值，会取出slice中后面的值，
		因为slice内存是连续的，所以可以取出数组后面的元素，
		通过cap函数可以看出slice的容量大小可知
	*/
	sc := sa[3:5]
	fmt.Println(string(sb), string(sc))

	s6 := make([]int, 3, 6)
	//打印地址
	fmt.Printf("%p\n", s6)
	s6 = append(s6, 1, 2, 3)
	//容量没有发生改变，所以内存地址没有发生变化
	fmt.Printf("%p, %v\n", s6, s6)
	s6 = append(s6, 1, 2, 3)
	//追加的元素超过容量，所以内存进行重新分配，地址也发生了变化
	fmt.Printf("%p %v\n", s6, s6)

	c := []int{1, 2, 3, 4, 5}
	s7 := c[2:5]
	s8 := c[1:3]
	fmt.Println(s7, s8)
	//slice指向的底层数组是同一个，修改其中一个其他的slice也会改变
	s7[0] = 6
	fmt.Println(s7, s8)

	d := []int{1, 2, 3, 4, 5}
	s9 := d[2:5]
	s10 := d[1:3]
	fmt.Println(s9, s10)
	//使用append追加元素超过原有slice容量后，内存重新分配
	s10 = append(s10, 1, 1, 12, 3, 1, 31, 13, 1, 13, 4, 35, 4)
	//此时对第一个slice进行修改，不会影响第二个slice，因为第二个slice内存地址以及不再指向数组d
	s9[0] = 6
	fmt.Println(s9, s10)

	//copy函数
	s11 := []int{1, 2, 3, 4, 5, 6}
	s12 := []int{7, 8, 9}
	copy(s11, s12)
	fmt.Println(s11, s12)
	s13 := []int{1, 2, 3, 4, 5, 6}
	s14 := []int{7, 8, 9}
	copy(s14, s13)
	fmt.Println(s13, s14)

	//指定下标复制
	copy(s14[2:4], s13[1:3])
}
