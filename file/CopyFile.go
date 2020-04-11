package main

import (
	"io"
	// "bufio"
	"fmt"
	"os"
)

func main() {
	//获取命令行参数
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFilename := list[1]
	dstFileName := list[2]
	if srcFilename == dstFileName {
		fmt.Println("源文件不能和目标文件相同")
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFilename)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建目标文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，往目标文件写，读多少写多少
	buf := make([]byte, 1024*4) //4k大小临时缓冲区

	for {
		//从源文件读取内容
		n, err := sF.Read(buf)
		if err != nil {
			//文件读取完毕
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}

		//往目标文件写，读多少写多少
		dF.Write(buf[:n])
	}
}
