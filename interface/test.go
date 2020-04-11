package main

import "fmt"

//定义接口
type USB interface {
	Name() string
	//嵌入接口
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

//实现接口方法
func (pc PhoneConnector) Name() string {
	return pc.name
}

//实现接口方法
func (pc PhoneConnector) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnector{"PhoneConnector"}
	a.Connect()
	//此处调用方法传递的是PhoneConnector类型，证明其确实实现了USB接口
	Disconnect(a)

	Disconnect1(a)

	b := PhoneConnector{"PhoneConnector"}
	b.Connect()
	//这里采用的赋值也是使用的拷贝
	c := Connector(b)
	//因为使用的是拷贝，所以这里的修改对c不生效
	b.name = "PC"
	c.Connect()

	//定义空接口
	var d interface{}
	fmt.Println(d == nil)
	var p *int = nil
	d = p
	fmt.Println(d == nil)
}

//显示定义方法，参数为接口类型
func Disconnect(usb USB) {
	//判断类型
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("Disconnected：", pc.Name())
		return
	}
	fmt.Println("Unknown device.")
}

//传递的参数类型为空接口，可以近似得认为空接口默认是所有接口的父类
func Disconnect1(usb interface{}) {
	//判断类型，直接由系统自动判断类型
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected：", v.Name())
	default:
		fmt.Println("Unknown device.")
	}

}
