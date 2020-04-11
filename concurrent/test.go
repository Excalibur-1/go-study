package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//并发执行
	//go Go()
	//time.Sleep(2 * time.Second)
	//创建channel
	c := make(chan bool)
	go func() {
		fmt.Println("Go2")
		//通道存值
		c <- true
	}()
	//取值，此处会阻塞获取通道中的值，通知main函数可以退出
	<-c
	A()
	fmt.Println("====================")
	//test()
	fmt.Println("====================")
	//test1()

	test3()

	test4()
}

func Go() {
	fmt.Println("Go1")
}

func A() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go3")
		//通道存值
		c <- true
		//关闭通道，如果不关闭，后续的循环操作会死循环
		close(c)
	}()
	//循环取值，当通道被关闭时结束循环
	for v := range c {
		fmt.Println(v)
	}
}

//通过通道实现并发执行
func test() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//创建通道，参数2为指定缓存个数
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		//并发执行
		go Go1(c, i)
	}
	//循环读取通道值
	for i := 0; i < 10; i++ {
		<-c
	}
}

//通过同步包实现并发执行
func test1() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		//并发执行
		go Go2(&wg, i)
	}
	wg.Wait()
}

func Go2(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}

func Go1(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func test3() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1")
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2")
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "Hi"
	c1 <- 3
	c2 <- "Hello"

	close(c1)
	close(c2)
	for i := 0; i < 2; i++ {
		<-o
	}
}

var c chan string

func Pingpang() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pingpang: Hi, #%d", i)
		i++
	}
}
func test4() {
	c = make(chan string)
	go Pingpang()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main: Hello, #%d", i)
		fmt.Println(<-c)
	}
}
