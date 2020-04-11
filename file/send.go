package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}

	defer f.Close()

	//使用切片建立缓冲
	buf := make([]byte, 1024*4)

	//循环读文件内容，读多少写多少
	for {
		n, err1 := f.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err1)
			}
			return
		}

		//发送内容
		conn.Write(buf[:n])
	}
}

func main() {
	fmt.Println("请输入文件：")

	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err = ", err1)
		return
	}

	defer conn.Close()

	//给接收方发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err = ", err2)
		return
	}

	//接收对方的回复，如果为ok，说明对方准备好，可以发送文件
	var n int
	buf := make([]byte, 1024)

	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err = ", err3)
		return
	}

	if "ok" == string(buf[:n]) {
		//发送文件内容
		SendFile(path, conn)
	}

}
