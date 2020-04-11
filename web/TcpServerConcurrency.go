package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完成自动关闭
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "addr connect successful")

	//新建切片，缓存数据
	buf := make([]byte, 2048)

	//循环读取
	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		//客户端发送时根据平台不同，会在后面追加字符，windows会追加"\r\n"，所以这里将n - 2
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, " exit")
			return
		}

		//将数据转换为大写返回给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//接收多个用户
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("err = ", err)
			return
		}

		//处理用户请求，新建协程
		go HandleConn(conn)
	}
}
