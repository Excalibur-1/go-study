package main

import "fmt"

func main() {
	//定义方式
	//底层数组指针为nil
	var s1 []int
	s2 := []byte{'a', 'b', 'c'}
	fmt.Println(s1)
	fmt.Println(s2)

	//使用make函数创建
	//创建长度为5，容量为2*5，初始值为0的切片
	slice1 := make([]int, 5)
	//创建长度为5，容量为7，初始值为0的切片
	slice2 := make([]int, 5, 7)
	//创建长度为5，容量为5，并已初始化的切片
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	//从数组创建slice，通过array[i:j]获取i表示数组下标开始位置，j表示结束位置，左闭右开
	//声明一个含有10个元素的元素类型为byte的数组
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	//声明两个含有byte的slice
	var a, b []byte
	//a指向数组的第三个元素开始到第五个元素结束，a含有的元素：arr[2],arr[3],arr[4]
	a = arr[2:5]
	//b是数组arr的另一个slice，b的元素是：arr[3],arr[4]
	b = arr[3:5]
	fmt.Println(a)
	fmt.Println(b)

	//从切片创建切片
	//oldSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	oldSlice := make([]int, 3, 6)
	//基于切片的前6个元素创建，没有的默认0，但是新切片容量必须小于等于原切片
	//oldSlice := make([]int, 3, 5)//会报错
	newSlice := oldSlice[:6]
	fmt.Println(newSlice)

	//切片操作
	//内置函数
	//返回切片长度
	fmt.Println(len(newSlice))
	//返回切片容量
	fmt.Println(cap(newSlice))
	//对切片追加元素，返回新的切片，不会影响原切片内容，且追加是在尾部追加
	fmt.Println(append(newSlice, 1))
	for _, v := range newSlice {
		fmt.Println(v)
	}
	//将slice3的数据拷贝到newSlice中，返回拷贝的个数
	fmt.Println(copy(newSlice, slice3))
	for _, v := range newSlice {
		fmt.Println(v)
	}
	//切片的容量与元素个数
	slice4 := make([]int, 5, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println(slice4)

	//切片操作
	//追加元素，容量未使用完的情况，将追加元素素个数，容量用完的情况下自动扩容，扩容是将原容量乘2
	slice1 = append(slice1, 1, 2)
	fmt.Println(slice1)

	//在一个切片上追加增加一个新的切片
	sliceTemp := make([]int, 3)
	slice1 = append(slice1, sliceTemp...)
	fmt.Println(slice1)
	fmt.Println("==========================")
	//切片拷贝
	s3 := []int{1, 3, 6, 9}
	//必须给够足够的空间，如果空间不够，则会丢弃部分数据
	s4 := make([]int, 10)
	//将s3拷贝到s4，注意不能写反，否则会出现意料之外的情况
	num := copy(s4, s3)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(num)

	//移除切片中的元素
	index := 2
	//采用的实际上是将切片划分成两个部分，然后根据开闭原则，将两部分合并起来，自然就移除了需要移除的元素
	s3 = append(s3[:index], s3[index+1:]...)
	fmt.Println(s3)

	//切片拷贝
	s5 := []int{1, 2, 3, 4, 5}
	s6 := []int{6, 7, 8}
	//复制s6前三个元素到s5前3位置，s5：[6,7,8,4,5]
	copy(s5, s6)
	//复制s5前三个元素到s6前3位置, s6:[6,7,8]
	copy(s6, s5)
	fmt.Println(s5)
	fmt.Println(s6)

	fmt.Println("================")

	//简单操作
	//slice的默认开始位置是0，slice[:n]等价于slice[0:n]
	//slice的第二个序列默认是切片的长度，slice[n:]等价于slice[n:len(slice)]
	//如果从一个数组里面直接获取slice，可以这样arr[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于arr[0:len(arr)]
	//切片的遍历可以使用for循环，也可以使用range函数

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	//演示
	// 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[:3]
	fmt.Println(aSlice)
	// 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[5:]
	fmt.Println(aSlice)
	// 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素
	aSlice = array[:]
	fmt.Println(aSlice)

	fmt.Println("=========================")
	//从slice获取slice
	// aSlice包含元素: d,e,f,g，len=4，cap=7
	aSlice = array[3:7]
	fmt.Println(aSlice)
	// bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[1:3]
	fmt.Println(bSlice)
	// bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[:3]
	fmt.Println(bSlice)
	//slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h，因aSlice的容量为7，实际上已经有数组中的数据在里面了
	bSlice = aSlice[0:5]
	fmt.Println(bSlice)
	// bSlice包含所有aSlice的元素: d,e,f,g
	bSlice = aSlice[:]
	fmt.Println(bSlice)
	fmt.Println("===================")

	//切片的截取，遵循左闭右开原则
	//s[n]：切片s中索引为位置为n的项
	//s[:]：从切片s的索引位置0到len(s)-1所获得的切片
	//s[low:]：从切片s的索引位置low到len(s)-1所获得的切片
	//s[:high]：从切片s的索引位置0到high所获得的切片
	//s[low:high]：从切片s的索引位置low到high所获得的切片
	//s[low:high:max]：从low到high的切片，且容量cap=max-low
	aSlice = array[:]
	cSlice := aSlice[3]
	fmt.Println(cSlice)
	dSlice := aSlice[:]
	fmt.Println(dSlice)
	eSlice := aSlice[5:]
	fmt.Println(eSlice)
	fSlice := aSlice[:5]
	fmt.Println(fSlice)
	gSlice := aSlice[5:7]
	fmt.Println(gSlice)
	hSlice := aSlice[5:7:10]
	fmt.Println(hSlice, cap(hSlice))

	//字符串转切片
	str := "hello,世界"
	//字符串转换为[]byte类型切片
	c := []byte(str)
	//字符串转换为[]rune类型切片
	d := []rune(str)
	fmt.Println(c, d)

	//切片存储结构,与数组相比，切片多了一个存储能力值的概念，即元素个数与分配空间可以是两个不同的值，其结构如下所示
	/*
		type slice struct {
			array = unsafe.Pointer		//指向底层数组的指针
			len int						//切片元素数量
			cap int						//底层数组的容量
		}
	*/
	/*
		所以切片通过内部的指针和相关属性引用数组片段，实现了变长方案，slice并不是真正意义上的动态数组。
		合理设置存储能力，可以大幅提升性能，比如知道最多元素个数为50，那么提前设置为50，而不是先设为30，
		可以明显减少重新分配内存的操作。
	*/

	//切片作为函数参数
	s7 := make([]int, 3)
	// 不会因为test函数内的append而改变
	fmt.Printf("main--%p\n", s7)
	test(s7)
	// 不会因为test函数内的append而改变
	fmt.Printf("main--%p\n", s7)
	fmt.Println("main--", s7)
}

//切片作为函数参数
func test(s []int) {
	//地址与main定义的slice一样
	fmt.Printf("test--%p\n", s)
	//append后生成了新的slice，没有影响原有的slice
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Printf("test--%p\n", s)
	fmt.Println("test--", s)
}
