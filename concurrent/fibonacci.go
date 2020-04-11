package main

import (
	"fmt"
)

//ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	//定义斐波那契数列起始值
	x, y := 1, 1
	for {
		//监听channel数据的流动
		select {
		//向管道写数据
		case ch <- x:
			x, y = y, x+y
		//从管道读取数据，读到就退出
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

func main() {
	//斐波那契数列：1,1,2,3,5,8,13...
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel读取内容
	go func() {
		for i := 0; i < 10; i++ {
			//从管道读取内容
			num := <-ch
			fmt.Println("num = ", num)
		}
		//读取完成，可以退出
		quit <- true
	}()

	//生产者，产生数字写入channel
	fibonacci(ch, quit)
}
