package main

import (
	"fmt"
	"time"
)

func main() {
	//定义有缓存channel，容量为3
	ch := make(chan int, 3)

	//len(ch)表示缓冲区剩余数据个数，cap(ch)表示缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程：i= ", i)

			//往channel写内容
			ch <- i
		}
	}()

	//延时
	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		//从管道中读取内容，没有内容时阻塞
		num := <-ch
		fmt.Println("num = ", num)
	}

}
