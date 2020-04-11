package main

import (
	"fmt"
	"time"
)

//生产者
func Producer(ch chan<- int) {
	i := 1
	for {
		ch <- i
		fmt.Println("Send:", i)
		i++
		time.Sleep(time.Second) //避免数据流动过快
	}
}

//消费者
func Consumer(ch <-chan int) {
	for {
		v := <-ch
		fmt.Println("Receive:", v)
		time.Sleep(time.Second * 2) //避免数据流动过快
	}
}

func main() {
	//生产消费模型中的缓冲区
	ch := make(chan int, 5)
	//启动生产者
	go Producer(ch)
	//启动消费者
	go Consumer(ch)

	//阻塞主程序退出
	for {

	}
	//当然，该模型也可以使用无缓冲模型，区别如下：
	//无缓冲生产消费模型：同步通信
	//有缓冲生产消费模型：异步通信
}
