package main

import "fmt"
import "time"

func main() {
	timer := time.NewTimer(3 * time.Second)

	//重置定时器
	timer.Reset(1 * time.Second)

	<-timer.C
	fmt.Println("时间到")
}

func main04() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程打印，定时器时间到")
	}()

	//停止定时器，导致子协程无法打印内容，因为定时器已经停止无法向管道写内容
	timer.Stop()

	for {

	}
}

func main03() {
	//定时2s，阻塞2s后产生一个事件，往channel写内容
	<-time.After(2 * time.Second)
	fmt.Println("时间到")
}

func main02() {
	timer := time.NewTimer(2 * time.Second)

	//验证timer.NewTimer(),时间到只会响应一次
	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

func main01() {
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("当前时间：", time.Now())

	//两秒后，往timer.C写数据，有数据后进行读取
	t := <-timer.C //channel没有数据前后阻塞
	fmt.Println("t = ", t)
}
