//文件练习
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	//创建文件
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)  //打印文件指针
	defer f.Close() //打开的资源在不使用时必须关闭

	/*
		使用Create()创建文件时：
		如果文件不存在，则创建文件。
		如果文件存在，则清空文件内内容。
		Create创建的文件任何人都可以读写。
	*/

	//打开文件，写入内容
	/*
		打开文件有两种方式：
		Open()：以只读的方式打开文件，若文件不存在则会打开失败
		OpenFile()：打开文件时，可以传入打开方式，该函数的三个参数：
		参数1：要打开的文件路径
		参数2：文件打开模式，如 O_RDONLY，O_WRONLY，O_RDWR，还可以通过管道符来指定文件不存在时创建文件
		参数3：文件创建时候的权限级别，在0-7之间，常用参数为6
	*/
	f, err = os.OpenFile("test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	fmt.Println(f.Name())

	/*
		常用的文件打开模式：
		O_RDONLY 	int = syscall.O_RDONLY		// 只读
		O_WRONLY	int = syscall.O_WRONLY		// 只写
		O_RDWR 		int = syscall.O_RDWR		// 读写
		O_APPEND 	int = syscall.O_APPEND		// 写操作时将数据追加到文件末尾
		O_CREATE 	int = syscall.O_CREATE		// 如果不存在则创建一个新文件
		O_EXCL 		int = syscall.O_EXCL		// 打开文件用于同步I/O
		O_TRUNC		int = syscall.O_TRUNC		// 如果可能，打开时清空文件
	*/

	//写文件
	//写入字节：Write()
	//写入文件内容，返回写入的字节数和错误（如果有）
	n, err := f.Write([]byte("hello"))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	fmt.Println("write number = ", n)
	//按字符串写：WriteString()
	//写入文件内容
	n, err = f.WriteString("hello")
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	fmt.Println("write number = ", n)
	//上述案例中，如果我们不想每次写入都会从头开始重新写入，那么需要将打开模式修改为：os.O_CREATE | os.O_WRONLY | os.O_APPEND

	//修改文件的读写指针位置：Seek()，包含两个参数
	/*
		参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
		参数2：偏移的开始位置，包括：
		io.SeekStart：从文件起始位置开始
		io.SeekCurrent：从文件当前位置开始
		io.SeekEnd：从文件末尾位置开始
	*/
	f, _ = os.OpenFile("test.txt", os.O_RDWR, 6)
	off, _ := f.Seek(5, io.SeekStart)
	fmt.Println(off)
	n, _ = f.WriteAt([]byte("111"), off)
	fmt.Println(n)
	defer f.Close()

	//获取文件描述信息：os.Stat()
	//Go的os包中定义了file类，封装了文件描述信息，同时也提供了Read、Write的实现。
	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println("stat err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	//获取到的fileInfo内部包含 文件名Name()、大小Size()、是否是目录IsDir() 等操作。

	//路径、目录操作
	//路径操作
	fmt.Println(filepath.IsAbs("./test.txt")) //false:判断是否是绝对路径
	path, err := filepath.Abs("./test.txt")
	if err != nil {
		fmt.Println("Abs err:", err)
		return
	}
	fmt.Println(path) //转换为绝对路径

	//创建目录，如果目录已存在，则会报错
	err = os.Mkdir("./test", os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Mkdir err:", err)
			return
		}
	}
	//创建多级目录，即使目录已存在也不会报错
	err = os.MkdirAll("./dd/rr", os.ModePerm)
	if err != nil {
		fmt.Println("MkdirAll err:", err)
		return
	}
	//贴士：OpenFile()可以用于打开目录。

	//删除文件
	/*err = os.Remove("test.txt")
	if err != nil {
		fmt.Println("remove err:", err)
		return
	}*/
	/*
		该函数也可用于删除目录（只能删除空目录）。
		如果要删除非空目录，需要使用 RemoveAll() 函数
	*/

}
