//go提供了一个channel（管道）数据类型，可以解决协程之间的通信问题！
//channel的本质是一个队列，遵循先进先出规则（FIFO），
//内部实现了同步，确保了并发安全！
package main

import (
	"fmt"
	"time"
)

func main() {
	//channel的语法：
	//channel在创建时，可以设置一个可选参数：缓冲区容量
	//创建有缓冲channel：make(chan int, 10)，创建一个缓冲长度为10的channel
	//创建无缓冲channel：make(chan int)，其实就是第二个参数为0
	//channel内可以存储多种数据类型，如下所示：
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})
	fmt.Println(ci, cs, cf)

	//从管道中读取，或者向管道写入数据，使用运算符：<-，他在channel的左边则是读取，右边则代表写入：
	ch1 := make(chan int)
	go func() {
		ch1 <- 10 //写入数据10
	}()
	num := <-ch1 //取出数据
	fmt.Println(num)
	//注意：无缓冲通道的收发操作必须在不同的两个goroutine间进行，因为通道的数据在没有接收方处理时，数据发送方会持续阻塞，所以通道的接收必定在另外一个 goroutine 中进行。
	//如果不按照该规则使用，则会引起经典的Golang错误fatal error: all goroutines are asleep - deadlock!:
	/*ch1 := make(chan int)
	ch1 <- 10    //写入数据10
	num := <-ch1 //取出数据
	fmt.Println(num)*/

	//无缓冲channel
	//无缓冲的channel是阻塞读写的，必须写端与读端同时存在，写入一个数据，则能读出一个数据：
	var write = func(ch chan int) {
		ch <- 100
		fmt.Printf("ch addr : %v\n", ch) //输出内存地址
		ch <- 200
		fmt.Printf("ch addr : %v\n", ch) //输出内存地址
		ch <- 300                        //该数据未读取，后续操作直接阻塞
		fmt.Printf("ch addr : %v\n", ch) //没有输出
	}

	var read = func(ch chan int) {
		//只读取两个数据
		fmt.Printf("取出的数据data1: %v\n", <-ch) //100
		fmt.Printf("取出的数据data2： %v\n", <-ch) //200
	}

	var ch chan int
	ch = make(chan int) //初始化

	//向协程写入数据
	go write(ch)
	//从协程读取数据
	go read(ch)

	//有缓冲channel
	//有缓冲的channel是非阻塞的，但是写满缓冲长度后，也会阻塞写入。
	var ch2 chan int        //声明一个有缓冲的channel
	ch2 = make(chan int, 2) //可以缓冲2个数据

	//向协程中写入数据
	go write(ch2) //第三个数据写入时会阻塞，因为缓冲只有2个

	//同样的，当数据全部读取完毕后，再次读取也会造成阻塞，如下所示：
	ch3 := make(chan int, 1)
	ch3 <- 11
	<-ch3
	//<-ch3
	//此时程序可以顺序运行，不会报错，这是与无缓冲通道的区别，但是当继续打开 注释 部分代码时，
	//通道阻塞，所有协程挂起，此时也会产生错误：fatal error: all goroutines are asleep - deadlock!。

	//纺织主协程提前退出，导致其他协程问完成任务
	//time.Sleep(time.Second * 3)

	//总结 无缓冲通道与有缓冲通道
	//无缓冲channel：
	//通道的容量为0，即 cap(ch) = 0
	//通道的个数为0，即 len(ch) = 0
	//可以让读、写两端具备并发同步的能力

	//有缓冲channel：
	//在make创建的时候设置非0的容量值
	//通道的个数为当前实际存储的数据个数
	//缓冲区具备数据存储的能力，到达存储上限后才会阻塞，相当于具备了异步的能力
	//有缓冲channel的阻塞产生条件：
	//	带缓冲通道被填满时，尝试再次发送数据会发生阻塞
	//	带缓冲通道为空时，尝试接收数据会发生阻塞

	//问题：为什么 Go语言对通道要限制长度而不提供无限长度的通道?
	//答：channel是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据 。
	//通道如果不限制长度，在生产速度大于消费速度时，内存将不断膨胀直到应用崩溃。

	//通道数据的遍历
	//channel只支持for-range的方式进行遍历
	ch4 := make(chan int)
	go func() {
		for i := 0; i <= 3; i++ {
			ch4 <- i
			//time.Sleep(time.Second)
		}
	}()
	for data := range ch4 {
		fmt.Println("data=", data)
		if data == 3 {
			break
		}
	}

	//通道关闭
	//通道是一个引用对象，支持GC回收，但是通道也可以被主动关闭：
	//ch5 := make(chan int)
	//close(ch) // 关闭通道
	//ch5 <- 1  //报错：send on closed channel

	//从通道中接收数据时，可以利用多返回值判断通道是否已经关闭：
	ch6 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch6 <- i
		}
		close(ch6)
	}()
	for {
		if num, ok := <-ch6; ok {
			fmt.Println("读到数据：", num)
		} else {
			fmt.Println(num)
			break
		}
	}
	//如果channel已经关闭，此时需要注意：
	//不能再向其写入数据，否则会引起错误：panic:send on closed channel
	//可以从已经关闭的channel读取数据，如果通道中没有数据，会读取通道存储的默认值

	//通道读写
	//默认情况下，管道的读写是双向的，但是为了对 channel 进行使用的限制，可以将 channel 声明为只读或只写:
	var chan1 chan<- int //声明只写channel
	var chan2 <-chan int //声明只读channel
	fmt.Println(chan1, chan2)
	//单向chanel不能转换为双向channel，但是双向channel可以隐式转换为任意类型的单向channel：
	// 只写端
	var onlyWrite = func(ch chan<- int) {
		ch <- 100
		fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
		ch <- 200
		fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
		ch <- 300                      // 该处数据未读取，后续操作直接阻塞
		fmt.Printf("ch addr：%v\n", ch) // 没有输出
	}
	// 只读端
	var onlyRead = func(ch <-chan int) {
		// 只读取两个数据
		fmt.Printf("取出的数据data1：%v\n", <-ch) // 100
		fmt.Printf("取出的数据data2：%v\n", <-ch) // 200
	}
	var ch7 chan int         //声明一个双向通道t
	ch7 = make(chan int, 10) //初始化
	//向通道写入数据
	go onlyWrite(ch7) //双向通道隐式转换为只写channel
	//从通道读取数据
	go onlyRead(ch7) //双向通道隐式转换为只读channel

	//显示转换：
	ch8 := make(chan int) //定义普通channel
	//ch9 := <-chan int (ch8)//转换为只读channel
	ch10 := chan<- int(ch8) //转换为只写channel
	go func() { ch10 <- 10 }()
	time.Sleep(time.Second * 3)

}
