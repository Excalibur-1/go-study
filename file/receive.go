package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func ReceiveFile(filName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(filName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	for {
		//接收对方发送过来的文件内容
		n, err1 := conn.Read(buf)
		if err1 != nil {
			if n == 0 || err1 == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err1)
			}

			return
		}

		//往文件写入内容
		f.Write(buf[:n])
	}
}

func main() {
	//监听网络请求
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err = ", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	//读取对方发送的文件名
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err = ", err2)
		return
	}

	filName := string(buf[:n])

	//回复ok，只能以切片形式写
	conn.Write([]byte("ok"))

	//接收文件内容
	ReceiveFile(filName, conn)
}
