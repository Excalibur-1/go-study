package main

import (
	"fmt"
	"time"
)

//创建channel
var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
	//向管道写数据
	ch <- 666
}

func person2() {
	//从管道取数据，如果没有数据会阻塞
	<-ch
	Printer("world")
}

func main() {
	go person1()
	go person2()

	for {

	}

}
