package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func slowFunc() {
	str := "hello world"
	for i := 0; i < 5; i++ {
		str += str
	}
}

func main() {
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("create cpu.prof err:", err)
		return
	}

	//获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("start cpu.prof err:", err)
		return
	}
	defer pprof.StopCPUProfile()

	//业务代码
	slowFunc()

	//获取内存相关信息
	f1, err := os.Create("mem.prof")
	defer f1.Close()
	if err != nil {
		fmt.Println("create mem.prof err:,", err)
		return
	}
	//runtime.GC() //是否获取最新的数据信息
	if err := pprof.WriteHeapProfile(f1); err != nil {
		fmt.Println("write cpu.prof err:", err)
		return
	}

	//获取协程相关信息
	f2, err := os.Create("goroutine")
	if err != nil {
		fmt.Println("Create goroutine.prof err:", err)
		return
	}
	if gProf := pprof.Lookup("goroutine"); gProf != nil {
		err = gProf.WriteTo(f2, 0)
	}

	//# 生成程序的二进制文件
	//go build -o program main.go				// 此时会按照代码中的要求生成多份prof文件
	//# 查看prof文件
	//go tool pprof program cpu.prof

	/*
		导入 "_ "net/http/pprof"包还可以实现以网页形式展示prof文件内容！
		程序执行前加上环境变量可以查看GC日志，如：GODEBUG=gctrace=1 go run main.go
	*/
}
