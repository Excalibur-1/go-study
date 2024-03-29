package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//使用完毕，需要关闭文件
	defer f.Close()

	var buf string

	for i := 0; i < 10; i++ {
		//"i = 1\n"，这个字符存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)
		fmt.Println("buf = ", buf)

		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("n = ", n)
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	//2K大小
	buf := make([]byte, 1024*2)

	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	//文件出错，同时没有读到结尾
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println(string(buf[:n]))
}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	for {
		//遇到'\n'结束读取，但是'\n'也读取进来了
		buf, err := r.ReadBytes('\n')
		if err != nil {
			//文件已经读完
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}

		fmt.Printf("buf = #%s#\n", string(buf))
	}

}

func main() {
	//关闭后，无法输出内容
	// os.Stdout.Close()
	//往标准输出设备写内容
	// fmt.Println("are you ok?")

	//标准设备文件(os.Stdout),默认已经打开，用户可以直接使用
	os.Stdout.WriteString("are you ok\n")

	//关闭后，无法输入内容
	// os.Stdin.Close()
	var a int
	fmt.Println("请输入a:")
	//从标准输入设备中读取内容放在a中
	fmt.Scan(&a)
	fmt.Println("a = ", a)

	path := "./demo.txt"
	// WriteFile(path)

	// ReadFile(path)

	ReadFileLine(path)
}
