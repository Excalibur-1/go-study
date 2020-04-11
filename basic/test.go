package main

import "fmt"
import "math"
import "strconv"

var i = 0

//批量定义变量测试
var (
	a = 1
	b = 2
)

type (
	byte     int8
	rune     int32
	byteSize int64
)

func main() {
	fmt.Println(i, a, b)

	fmt.Println(math.MinInt32)

	var a [1]byte
	fmt.Println(a)

	var b byte
	var c rune
	var d byteSize
	fmt.Println(b, c, d)

	var e int = 1
	var f = 1
	g := 1
	h := false
	fmt.Println(e, f, g, h)

	var j int = 65
	var k string
	k = string(j)
	fmt.Println(k)
	//数字转字符串
	k = strconv.Itoa(j)
	//字符转数字
	j, _ = strconv.Atoi(k)
	fmt.Println(k)
	fmt.Println(j)
}
