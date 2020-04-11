package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	//文件读取
	//文件读取的接口位于io包，file文件类是这些接口的实现类
	//直接读取read()，read()实现的是按字节数读取
	/*readByte := make([]byte, 128) // 指定要读取的长度
	for {
		//将数据读取入切片，返回值 n 是实际读取到的字节数
		n, err := f.Read(readByte)
		//错误不为空，或者读到了文件末尾，EOF即end of file
		if err != nil && err != io.EOF {
			fmt.Println("read file err:", err)
			break
		}

		fmt.Println("read:", string(readByte[:n]))
		//如果n < 128表示文件已经读完了
		if n < 128 {
			fmt.Println("read end")
			break
		}
	}*/

	/*
		bufio的写操作
		bufio封装了io.Reader、io.Writer接口对象，并创建了另一个也实现了该接口的对象：bufio.Reader、bufio.Writer。通过该实现，bufio实现了文件的缓冲区设计，可以大大提高文件I/O的效率。
		使用bufio读取文件时，先将数据读入内存的缓冲区（缓冲区一般比要比程序中设置的文件接收对象要大），这样就可以有效降低直接I/O的次数。
		bufio.Read([]byte)相当于读取大小len(p)的内容：
		当缓冲区有内容时，将缓冲区内容全部填入p并清空缓冲区
		当缓冲区没有内容且len(p)>len(buf)，即要读取的内容比缓冲区还要大，直接去文件读取即可
		当缓冲区没有内容且len(p)<len(buf)，即要读取的内容比缓冲区小，读取文件内容并填满缓冲区，并将p填满
		以后再次读取时，缓冲区有内容，将缓冲区内容全部填入p并清空缓冲区（和第一步一致）
	*/
	//创建读对象
	reader := bufio.NewReader(f)

	//读一行数据
	byt, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("ReadBytes err:", err)
	}
	fmt.Println(string(byt))

	//ReadString() 函数也具有同样的功能，且能直接读取到字符串数据，无需转换，示例：读取大文件的全部数据
	for {
		str, err := reader.ReadString('\n') //从缓冲区读取：读取到特定字符结束，按行读取
		if err != nil && err != io.EOF {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("str = ", str)
		if err == io.EOF {
			fmt.Println("read end")
			break
		}
	}

	//在Unix设计思想中，一切皆文件，命令行输入也可以作为文件读入：
	/*reader = bufio.NewReader(os.Stdin)
	//假设命令行以 - 开始
	s, err := reader.ReadString('-')
	if err != nil {
		fmt.Println("ReadString err:", err)
		return
	}
	fmt.Println(s)*/

	//io/ioutil包读取文件
	//ioutil直接读取文件
	ret, err := ioutil.ReadFile("test1.txt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	fmt.Println(string(ret))

	//bufio的写操作
	writer := bufio.NewWriter(f)
	n, err := writer.WriteString("hello world")
	if err != nil {
		fmt.Println("WriterString err:", err)
		return
	}
	// 必须刷新缓冲区：将缓冲区的内容写入文件中。如果不刷新，则只会在内容超出缓冲区大小时写入
	err = writer.Flush()
	if err != nil {
		fmt.Println("Flush err:", err)
		return
	}
	fmt.Println(n)

	//io/ioutil的写操作
	s := "你好世界"
	//文件不存在时，会自动创建文件
	err = ioutil.WriteFile("test2.txt", []byte(s), os.ModePerm)
	if err != nil {
		fmt.Println("WriteFile err:", err)
		return
	}

	f1, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f1.Close()
	//读取前五个字节，假设读取的文件内容为：hello world!
	//创建一个字节的切片
	bs := []byte{0}
	_, err = f1.Read(bs)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Println("读到的数据是：", string(bs)) //h

	//移动光标
	//光标从开始位置（h之前），移动4位，到达o之前
	_, err = f1.Seek(4, io.SeekStart)
	if err != nil {
		fmt.Println("seek err:", err)
		return
	}
	_, err = f1.Read(bs)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Println("读取到的数据是：", string(bs)) //o
	/*
		通过记录光标的位置，可以实现断点续传：假设已经下载了1KB文件，
		即本地临时文件存储了1KB，此时断电，重启后通过本地文件大小、Seek()方法获取到上次读取文件的光标位置即可实现继续下载！
	*/
}
