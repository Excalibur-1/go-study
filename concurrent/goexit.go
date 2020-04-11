package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer fmt.Println("ccccccccc")

	//终止此函数
	// return

	//终止所在的协程
	runtime.Goexit()

	fmt.Println("dddddddddddddd")
}

func main() {

	go func() {
		fmt.Println("aaaaaaaaa")

		test()

		fmt.Println("bbbbbbbbbb")
	}()

	//runtime.Goexit()：用于立即终止当前协程运行，调度器会确保所有已注册defer延迟调用被执行
	for i := 1; i <= 5; i++ {
		defer fmt.Println("defer ", i)
		go func(i int) {
			if i == 3 {
				runtime.Goexit()
			}
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)

	//死循环，不让主协程退出
	for {

	}

}
