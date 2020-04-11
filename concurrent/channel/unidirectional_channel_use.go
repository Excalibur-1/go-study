package main

import "fmt"

//生产者，只能写，不能读
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	//写完后记得关闭通道
	close(out)
}

//消费者，只能读，不能写
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}

func main() {
	//定义channel
	ch := make(chan int)

	//生产者
	go producer(ch)

	//消费者
	consumer(ch)
}
