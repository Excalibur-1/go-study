package main

import (
	"fmt"
	"time"
)

func main() {
	/*for i := 1; i <= 10; i++ {
		go func(){
			fmt.Println(i)		// 全部打印11：因为开启协程也会耗时，协程没有准备好，循环已经走完
		}()
	}
	time.Sleep(time.Second)*/

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i) // 打印无规律数字
		}(i)
	}
	time.Sleep(time.Second)
}
