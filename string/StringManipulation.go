package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//"hello go"是否包含"hello"
	fmt.Println(strings.Contains("hello go", "hello"))

	buf := "hello @abc @mike"
	s := strings.Split(buf, "@")
	fmt.Println("s = ", s)

	s = strings.Fields(buf)
	fmt.Println("s = ", s)

	//转换字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个参数为要追加的数，第三个参数为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abchellogo")

	fmt.Println("slice = ", string(slice))

	//其他类型转换为字符串
	str := strconv.FormatBool(false)
	fmt.Printf("str type is %T, str = %s", str, str)
}
