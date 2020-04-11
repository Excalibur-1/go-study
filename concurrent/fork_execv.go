//演示fork execv系统函数调用
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("当前进程id：", os.Getpid())

	procAttr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	//开启一个子进程，调用系统的echo函数输出:hello,world!
	process, err := os.StartProcess("/bin/echo", []string{"", "hello,world!"}, procAttr)
	if err != nil {
		fmt.Println("进程启动失败：", err)
		os.Exit(2)
	} else {
		fmt.Println("子进程id：", process.Pid)
	}
	time.Sleep(time.Second)
	//根据该方式，就可以很容运行计算机上的其他任何程序，包括自身的命令行、Java程序等等。
}
