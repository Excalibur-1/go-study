//Go语言的接口在命名时，一般会在单词后面添加er，如写操作的接口叫做Writer
//当方法名首字母大写，且实现的接口首字母也是大写，则该方法可以被接口所在包之外的代码访问
//方法与接口中的方法签名一致（方法名、参数列表、返回列表都必须一致）
//参数列表和返回值列表中的变量名可以被忽略，如：type writer interfae{ Write([]byte) error}
//接口中所有的方法都必须被实现，如果不实现，编译器不会显示提示错误，编译时会报错，
//如果编译时发现实现接口的方法签名不一致，则会报错：does not implement。
package main

import "fmt"

//声明格式：
/*type 接口类型名 interface {
	方法名1(参数列表) 返回值列表
	方法名2(参数列表) 返回值列表
	...
}
*/

//运输方式接口
type Transporter interface {
	BicycleTran()
	CarTran()
}

//驾驶员
type Driver struct {
	Name string
	Age  int
}

//实现Transporter接口的BicycleTran()方法
func (d *Driver) BicycleTran() {
	fmt.Println("使用自行车运输")
}

//实现Transporter接口的CarTran()方法
func (d *Driver) CarTran() {
	fmt.Println("使用小汽车运输")
}

//只要实现了Transporter接口的类型都可以作为参数
func trans(t Transporter) {
	t.BicycleTran()
}

//go接口的特点
//在上述示例中，Go无须像Java那样显式声明实现了哪个接口，即为非侵入式，接口编写者无需知道接口被哪些类型实现，接口实现者只需要知道实现的是什么样子的接口，但无需指明实现了哪个接口。编译器知道最终编译时使用哪个类型实现哪个接口，或者接口应该由谁来实现。
//类型和接口之间有一对多和多对一的关系，即：
//一个类型可以实现多个接口，接口间是彼此独立的，互相不知道对方的实现
//多个类型也可以实现相同的接口。

//游戏服务
type Service interface {
	Start()
	Log(string)
}

//日志器
type Logger struct {
}

//日志输出方法
func (g *Logger) Log(s string) {
	fmt.Println("日志：", s)
}

//游戏服务
type GameService struct {
	Logger
}

//实现游戏服务的Start方法
func (g *GameService) Start() {
	fmt.Println("游戏服务启动")
}

//接口嵌套
//Go中不仅结构体之间可以嵌套，接口之间也可以嵌套。接口与接口嵌套形成了新的接口，
//只要接口的所有方法被实现，则这个接口中所有嵌套接口的方法均可以被调用。
//定义一个写接口
type Writer interface {
	Write([]byte) (int, error)
}

//定义一个读接口
type Reader interface {
	Read() error
}

//定义一个嵌套接口
type IO interface {
	Writer
	Reader
}

//空接口
//空接口定义
//空接口是接口的特殊形式，没有任何方法，因此任何具体的类型都可以认为实现了空接口。
var any interface{}

//空接口作为参数参考
func Test(i interface{}) {
	fmt.Printf("%T\n", i)
}

//定义一个通用接口：动物接口
type Animal interface {
	Breath()
}

type Flyer interface {
	Fly()
}

type Swimer interface {
	Swim()
}

//定义一个鸟类
type Bird struct {
	Name string
	Food string
	Kind string
}

func (b *Bird) Breath() {
	fmt.Println("鸟在陆地呼吸")
}

func (b *Bird) Fly() {
	fmt.Printf("%s在飞\n", b.Name)
}

//定义一个鱼类
type Fish struct {
	Name string
	Kind string
}

func (f *Fish) Breath() {
	fmt.Println("鱼在水下呼吸")
}

func (f *Fish) Swim() {
	fmt.Printf("%s在游泳\n", f.Name)
}

/*func Display(a Animal) {
	//a.Breath()
	//调用实现类的成员：此时会报错
	//fmt.Printf(a.Name)

	//接口类型无法直接访问其实现类的成员，需要使用断言，对接口的类型进行判断，类型断言格式：
	//t := i.(T)//不安全写法：如果i没有完全实现T接口的方法，这个语句将会触发宕机
	//t, ok := i.(T)// 安全写法：如果接口未实现接口，将会把ok掷为false，t掷为T类型的0值
	//上述案例可以书写为：
	//直接调用接口的方法
	a.Breath()
	//注意：这里必须是 *Bird类型，因为是*Bird实现了接口，不是Bird实现了接口
	instance, ok := a.(*Bird)
	if ok {
		//得到具体的实现类，才能访问实现类的成员
		fmt.Println("该鸟类的名字是：", instance.Name)
	} else {
		fmt.Println("该动物不是鸟类")
	}
}*/

//接口类型转换
//在接口定义时，其类型已经确定，因为接口的本质是方法签名的集合，
//如果两个接口的方法签名结合相同（顺序可以不同），则这2个接口之间不需要强制类型转换就可以相互赋值，
//因为go编译器在校验接口是否能赋值时，比较的是二者的方法集。
//在上面，函数Display接收的是Animal接口类型，在断言后转换为了别的类型：*Bird(实现类指针类型)
//其实，断言还可以将接口转换成另外一个接口：
func Display(a Animal) {
	instance, ok := a.(Flyer)
	if ok {
		instance.Fly()
	} else {
		fmt.Println("该动物不会飞")
	}
}

//New出Animal的函数
func NewAnimal(kind string) Animal {
	switch kind {
	case "鸟类":
		return &Bird{}
	case "鱼类":
		return &Fish{}
	default:
		return nil
	}
}

func main() {
	d := &Driver{
		Name: "张三",
		Age:  27,
	}
	trans(d)

	//在下面案例中，即使没有接口也能运行，但是当存在接口时，会隐式实现接口，让接口给类提供约束。
	//使用接口调用了结构体中的方法，也可以理解为实现了面向对象中的多态。
	s := new(GameService)
	s.Start()
	s.Log("hello")

	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)

	Test(3)
	Test("hello")

	//利用空接口，可以实现任意类型的存储
	m := make(map[string]interface{})
	m["name"] = "李四"
	m["age"] = 30
	fmt.Println(m)

	//从空接口获取值
	//保存到空接口的值，如果直接取出指定类型的值时，会发生编译错误：
	var a int = 1
	var i interface{} = a
	//var b int = i//这里编译报错（类型不一致），可以采用：b := i 的形式
	b := i
	fmt.Printf("%T\n", b)

	//空接口值比较
	//类型不同的空接口比较
	var c interface{} = 100
	var f interface{} = "hi"
	fmt.Println(c, f, c == f) //false

	//不能比较空接口中的动态值
	var e interface{} = []int{10}
	var g interface{} = []int{20}
	fmt.Println(e, g)
	//fmt.Println(e == g) //运行报错
	/*
		map		不可比较，会发生宕机错误
		切片		不可比较，会发生宕机错误
		通道		可比较，必须由同一个make生成，即同一个通道才是true
		数组		可比较，编译期即可知道是否一致
		结构体	可比较，可诸葛比较结构体的值
		函数		可比较
	*/

	var bird = &Bird{
		"斑鸠",
		"蚂蚱",
		"鸟类",
	}
	Display(bird)

	//一个实现类往往实现了很多接口，为了精准类型查询，可以使用switch语句来判断对象类型：
	/*var v1 interfaceP{} = ...
	switch v := v1.(type) {
	case int:
	case string:
		...
	}*/

	//多态
	//获取的是动物接口类型，但是实现类是鸟类
	a1 := NewAnimal("鸟类")
	a1.Breath() //鸟在陆地呼吸

	//获取的是动物接口类型，但是实现类是鱼类
	a2 := NewAnimal("鱼类")
	a2.Breath() //鱼在水下呼吸

	//演示方法接受者使用指针和非指针的区别
	inter1 := newStruct1(2)
	inter1.squareValue()
	inter1.doubleValue()

}

type interface1 interface {
	doubleValue()
	squareValue()
}

func newStruct1(value int) interface1 {
	return &struct1{value}
}

type struct1 struct {
	value int
}

func (s struct1) doubleValue() {
	s.value = s.value * 2
	fmt.Println(s.value)
}

func (s *struct1) squareValue() {
	s.value = s.value * s.value
	fmt.Println(s.value)
}
