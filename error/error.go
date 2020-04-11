//
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
)

//defer最佳实践
//案例一：defer处理资源

//没有使用defer时打开文件处理代码：
func OpenFile(file string) int {
	f, err := os.Open(file)
	if err != nil {
		return 0
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return 0
	}
	fmt.Println(info)
	f.Close()
	return 0
}

//使用defer优化：
func OpenFile2(file string) int {
	f, err := os.Open(file)
	if err != nil {
		return 0
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		// f.Close()			//这句已经不需要了
		return 0
	}
	fmt.Println(info)

	//后续一系列文件操作后执行关闭
	// f.Close()			//这句已经不需要了
	return 0
}

//案例二：并发使用map的函数。
//无defer代码：
/*var (
	mutex   sync.Mutex
	testMap = make(map[string]int)
)

func getMapValue(key string) int {
	mutex.Lock() //对共享资源加锁
	value := testMap[key]
	mutex.Unlock()
	return value
}
*/
//上述案例是很常见的对并发map执行加锁执行的安全操作，使用defer可以对上述语义进行简化：
var (
	mutex   sync.Mutex
	testMap = make(map[string]int)
)

func getMapValue(key string) int {
	mutex.Lock() //对共享资源加锁
	defer mutex.Unlock()
	return testMap[key]
}

//defer无法处理全局资源
//使用defer语句, 可以方便地组合函数/闭包和资源对象，即使panic时，defer也能保证资源的正确释放。
//但是上述案例都是在局部使用和释放资源，如果资源的生命周期很长， 而且可能被多个模块共享和随意传递的话，defer语句就不好处理了。
//Go的runtime包的func SetFinalize(x, f interface{})函数可以提供类似C++析构函数的机制。
//示例：包装一个文件对象，在没有人使用的时候能够自动关闭。
type MyFile struct {
	f *os.File
}

func NewFile(name string) (*MyFile, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	runtime.SetFinalizer(file, file.Close)
	return &MyFile{f: file}, nil
}

//返回错误前，需要定义会产生哪些可能的错误，在Go中，使用errors包进行错误的定义，格式如下：
//var err = errors.New("发生了错误")
//提示：错误字符串相对固定，一般在包作用于声明，应尽量减少在使用时直接使用errors.New返回。
//errors.New使用示例:
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// implementation
	return 0, nil
}

//实现错误接口案例：
//声明一种解析错误
type ParseError struct {
	FileName string
	Line     int
}

//实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.FileName, e.Line)
}

//创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

/*
recover 宕机恢复
Go提供的recover机制：由运行时抛出，或者开发者主动触发的panic，可以使用defer和recover实现错误捕捉和处理，让代码在发生崩溃后允许继续执行。
在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过try/catch机制捕获异常，没有被捕获的严重异常会导致宕机，
捕获的异常可以被忽略，让代码继续执行。Go没有异常系统，使用panic触发宕机类似于其他语言的抛出异常，recover的宕机恢复机制就对应try/catch机制。
panic和defer的组合：
有panic没有recover，程序宕机
有panic也有recover，程序不会宕机，执行完对应的defer后，从宕机点退出当前函数后继续执行
*/
func test(num1 int, num2 int) {
	defer func() {
		err := recover() // recover内置函数，可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	fmt.Println(num1 / num2)
}

func main() {
	var e error
	e = newParseError("main.go", 1)
	fmt.Println(e.Error())
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line:%d \n", detail.FileName, detail.Line)
	default:
		fmt.Println("other error")
	}

	test(2, 0)
	fmt.Println("after...") // 该句会输出
}
