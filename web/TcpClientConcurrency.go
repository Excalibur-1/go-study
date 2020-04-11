package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer conn.Close()

	go func() {
		str := make([]byte, 1024)
		//从键盘输入内容，发送给服务器
		for {
			n, err2 := os.Stdin.Read(str)
			if err2 != nil {
				fmt.Println("os.Stdin.Read err = ", err2)
				return
			}

			//发送输入的内容给服务器
			conn.Write(str[:n])
		}
	}()

	//切片缓冲
	buf := make([]byte, 1024)

	for {
		//接收服务器请求
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("conn.Read err = ", err1)
			return
		}
		//打印接收到的内容
		fmt.Println(string(buf[:n]))
	}

}
