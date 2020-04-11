package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//指定以1核运算
	// runtime.GOMAXPROCS(1)
	//指定以4核运算
	runtime.GOMAXPROCS(4)
	//获取当前系统核心数
	fmt.Println(runtime.NumCPU())
	for {
		go fmt.Print(1)
		fmt.Print(0)
		time.Sleep(time.Second)
	}

	//贴士：在Go1.5之后，程序已经默认运行在多核上，无需上述设置

}
