package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigRecv := make(chan os.Signal, 1)                   //创建接收通道
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT} //创建信号类型
	signal.Notify(sigRecv, sigs...)
	for sig := range sigRecv { // 循环接收通道中的信号，通道关闭后，for会立即停止
		fmt.Println(sig)
	}

}
