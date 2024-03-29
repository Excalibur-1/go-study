package main

import (
	"bytes"
	"fmt"
)

//函数声明格式
/*
函数名首字母小写为私有，大写为公有；
参数列表可以有0-多个，多参数使用逗号分隔，不支持默认参数；
返回值列表返回值类型可以不用写变量名
如果只有一个返回值且不声明类型，可以省略返回值列表与括号
如果有返回值，函数内必须有return
*/
/*func 函数名(参数列表)(返回值列表){
	//函数体
	return 返回值列表
}
*/

//函数常见写法

//无返回值
func fn() {

}

//go推荐给函数返回值起一个变量名
func fn1() (result int) {
	return 1
}

//第二种返回值写法
func fn2() (result int) {
	result = 1
	return
}

//多返回值情况
func fn3() (int, int, int) {
	return 1, 2, 3
}

//go多返回值推荐写法，多个参数类型如果相同，可以简写为： a,b,c int 这种形式
func fn4() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return
}

/*
值传递和引用传递
不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是地址的拷贝，
一般来说，地址拷贝效率高，因为数据量小，而值拷贝决定拷贝的 数据大小，数据越大，效率越低。
如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。
*/

//可变参数
func joinStrings(sList ...string) string {
	var buf bytes.Buffer
	for _, s := range sList {
		buf.WriteString(s)
	}
	return buf.String()
}

//参数传递
//实际打印函数
func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

//封装打印函数
func print(sList ...interface{}) {
	//将sList可变参数切片完整传递给下一个函数
	rawPrint(sList...)
}

//函数类型
/*
函数去掉函数名、参数名和{}后的结果即是函数类型，可以使用%T打印该结果。
两个函数类型相同的前提是：拥有相同的形参列表和返回值列表，且列表元素的次序、类型都相同，形参名可以不同。
*/
func mathSum(a, b int) int {
	return a + b
}

func mathSub(a, b int) int {
	return a - b
}

//定义一个函数类型
type MyMath func(int, int) int

//定义的函数类型作为参数使用
func Test(f MyMath, a, b int) int {
	return f(a, b)
}

//特殊函数
/*
Go语言中，除了可以在全局声明中初始化实体，也可以在init函数中初始化。
init函数是一个特殊的函数，它会在包完成初始化后自动执行，执行优先级高于main函数，
并且不能手动调用init函数，每一个文件有且仅有一个init函数，
初始化过程会根据包的以来关系顺序单线程执行。
*/
func init() {
	//在这里可以书写一些初始化操作
	fmt.Println("init...")
}

func main() {
	fmt.Println(joinStrings("pig", " and", " bird"))

	print(1, 2, 3)

	//匿名函数
	/*
	   匿名函数可以看做函数字面量，所有直接使用函数类型变量的地方都可以由匿名函数代替。
	   匿名函数可以直接赋值给函数变量，可以当做实参，也可以作为返回值使用，还可以直接被调用。
	*/
	a := 3
	//f1即为匿名函数
	f1 := func(num int) {
		//匿名函数访问外部变量
		fmt.Println(num)
	}
	f1(a)

	//匿名函数自调
	func() {
		fmt.Println(a)
	}()

	//取最大值最小值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Printf("max=%d, min=%d\n", x, y)

	//调用函数参数为函数类型的函数，此处有点类似多态的运用
	sum := Test(mathSum, 1, 1)
	fmt.Println(sum)

	//通常可以把函数类型当做一种引用类型，实际函数类型变量和函数名都可以当做指针变量，指向函数代码开始的位置，没有初始化的函数默认值是nil。

	/*
		Go函数特性总结:
		支持有名称的返回值；
		不支持默认值参数；
		不支持重载；
		不支持命名函数嵌套，匿名函数可以嵌套；
		Go函数从实参到形参的传递永远是值拷贝，有时函数调用后实参指向的值发生了变化，是因为参数传递的是指针的拷贝，实参是一个指针变量，传递给形参的是这个指针变量的副本，实质上仍然是值拷贝；
		Go函数支持不定参数；
	*/

	//new函数
	//new函数可以用来创建变量。表达式new(T)将创建一个T类型的匿名变量，
	//初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T：
	//p为*int类型，指向匿名的int变量
	p := new(int)
	fmt.Println(*p)
	//设置int匿名变量值为2
	*p = 2
	fmt.Println(*p)

	//new函数还可以用来为结构体创建实例
	type file struct {
	}
	f := new(file)
	fmt.Println(f)

	fmt.Println(newInt1())
	fmt.Println(newInt2())

	//new只是一个预定义函数，并不是一个关键字，所以new也有可能会被项目定义为别的类型。

	//make函数
	//make函数经常用来创建切片、map、管道
	m1 := map[string]int{}
	m2 := make(map[string]int, 10)
	fmt.Println(m1, m2)
	//上面展示了两种map的创建方式，其不同点是第一种创建方式无法预估长度，当长度超过了当前长度时，会引起内存的拷贝！！
	//第二种创建方式直接限定了长度，这样能有效提升性能！

}

//new函数其实是语法糖，不是新概念，如下所示的两个函数其实拥有相同的行为。
func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
