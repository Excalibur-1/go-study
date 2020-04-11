//Go语言从语言层面原生提供了协程支持，即 goroutine，执行goroutine只需极少的栈内存(大概是4~5KB)，所以Go可以轻松的运行多个并发任务。
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick:", times)
		time.Sleep(time.Second)
	}
}

func main() {
	//Go中以关键字 go 开启协程：
	//go say("Go")                //以协程方式执行say函数
	//say("noGo")                 //以普通方式执行say函数
	//time.Sleep(time.Second * 3) //睡眠5秒：纺织协程未执行完毕，主程序退出

	//命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。两个环节可以同时进行，
	//直到按 Enter键时将输入的内容写入 input变量中井返回， 整个程序终止。
	go running()
	var input string
	n, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Scanln err:", err)
	}
	fmt.Println(n)

}
