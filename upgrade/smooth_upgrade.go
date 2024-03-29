/*
平滑升级
服务器在升级时，正在处理的请求需要等待其完成，再退出。Go1.8之后支持该设计。

实现步骤原理：

1 fork一个子进程，继承父进程的监听socket
2 子进程启动后，接收新的连接，父进程处理原有请求并且不再接收新请求
当系统重启或者升级时，正在处理的请求以及新来的请求该如何处理？

正在处理的请求如何处理：

等待处理完成之后，再推出，Go1.8之后已经支持。比如每来一个请求，计数+1，处理完一个请求，计数-1，当计数为0时，则执行系统升级。

新进来的请求如何处理：

Fork一个子进程，继承父进程的监听socket（os.Cmd对象中的ExtraFiles参数进行传递，并继承文件句柄）
子进程启动成功后，接收新的连接
父进程停止接收新的连接，等已有的请求处理完毕，退出，优雅重启成功。
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var child *bool

func startChild(file *os.File) {
	args := []string{"-child"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.ExtraFiles = []*os.File{file}
	err := cmd.Start()
	if err != nil {
		fmt.Printf("start child failed err:%v\n", err)
		return
	}
}

func init() {
	//命令行有child选项，则是子进程，没有则是父进程
	child = flag.Bool("child", false, "继承于父进程")
	flag.Parse()
}

func readFromParent() {
	//fd=0 标准输出，=1标准输入，=2标准错误输出，=3 ExtraFiles[0]，=4 ExtraFiles[1]
	f := os.NewFile(3, "")
	count := 0
	for {
		str := fmt.Sprintf("hello, i'child process,write:%d line \n", count)
		count++
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Printf("wrote string failed,err:%v\n", err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
	}
}

func main() {
	if child != nil && *child == true {
		fmt.Printf("继承于父进程的文件句柄\n")
		readFromParent()
		return
	}

	//父进程逻辑
	file, err := os.OpenFile("./test_inherit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	_, err = file.WriteString("parent write one line \n")
	if err != nil {
		fmt.Printf("parent write failed,err:%v\n", err)
		return
	}

	startChild(file)
	fmt.Println("parent exited")
}
