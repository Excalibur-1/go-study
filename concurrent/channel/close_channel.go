package main

import (
	"fmt"
)

func main() {
	//定义无缓存channel
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			//往channel写内容
			ch <- i
		}
		//关闭管道
		close(ch)
	}()

	//管道关闭时自动跳出循环
	for num := range ch {
		fmt.Println("num = ", num)
	}
}

func main01() {
	//定义无缓存channel
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			//往channel写内容
			ch <- i
		}
		//关闭管道
		close(ch)
	}()

	for {
		//管道关闭时ok返回false
		if num, ok := <-ch; ok {
			fmt.Println("num = ", num)
		} else {
			break
		}
	}

}
