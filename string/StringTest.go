package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "你好"

	//输出6
	fmt.Println(len(str))

	//输出2，因为中文占3个字节，所以需要使用下面的方法进行统计字符数
	fmt.Println(utf8.RuneCountInString(str))

	str = "hello"

	length := len(str)
	for i := 0; i < length; i++ {
		fmt.Printf("%d %v %c\n", i, str[i], str[i])
	}

	str = "你好"
	//中文字符串遍历只能使用range形式
	for i, ch := range str {
		fmt.Printf("%d %c\n", i, ch)
	}

	//类型转换
	num := 12
	//输出转换后的数据类型
	fmt.Printf("%T\n", string(num))

	//字符串连接
	var stringBuilder bytes.Buffer

	stringBuilder.WriteString(str)
	stringBuilder.WriteString("，世界！")

	//数据拼接后的字符串
	fmt.Println(stringBuilder.String())

	//索引查找没找到则返回-1，中文是按3个字节来算的，故“好”字返回3
	index := strings.Index(stringBuilder.String(), "好")
	fmt.Println("index:", index)

	index = strings.Index("hello", "l")
	//返回2
	fmt.Println("index:", index)

	contains := strings.Contains(str, "你")
	//返回true
	fmt.Println("contains:", contains)

	//用指定字符串连接切片中的字符串
	joinString := strings.Join([]string{"1", "2", "3"}, str)
	//输出：1你好2你好3
	fmt.Println("joinString:", joinString)

	//替换指定字符串中的字符串为新字符串，n 参数小于0表示替换次数不限
	replace := strings.Replace(str, "你", "you", -1)
	fmt.Println("replace:", replace)

	//按指定字符串分割字符串，返回字符串切片
	split := strings.Split(stringBuilder.String(), "，")
	for index, str := range split {
		fmt.Printf("split index:%d, str:%s\n", index, str)
	}

	//取出字符串头尾指定的字符串
	trim := strings.Trim(str, "你")
	fmt.Println("trim:", trim)

	//去除字符串中的空格，返回按空格分割的切片
	fields := strings.Fields("你好， 世界 ！")
	for index, str := range fields {
		fmt.Printf("fileds[%d]:[%s]\n", index, str)
	}

	//字符串转换函数
	//Append系列函数将整数等转换为字符后，添加到现有的字节切片中
	str1 := make([]byte, 0, 100)
	str1 = strconv.AppendInt(str1, 4567, 10)
	str1 = strconv.AppendBool(str1, false)
	str1 = strconv.AppendQuote(str1, "abcefg")
	str1 = strconv.AppendQuoteRune(str1, '单')
	for _, v := range str1 {
		//输出的是对应字符在Ascii码表的10进制表示数值
		fmt.Printf("%v, ", v)
	}
	fmt.Println()
	fmt.Println(string(str1))

	//Format系列函数把其他类型的数据转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1024)
	fmt.Println(a, b, c, d, e)

	//Parse系列函数把字符串转换为对应其他类型
	f, _ := strconv.ParseBool("false")
	g, _ := strconv.ParseFloat("123.23", 64)
	h, _ := strconv.ParseInt("1234", 10, 64)
	i, _ := strconv.ParseUint("12345", 10, 64)
	j, _ := strconv.Atoi("1024")
	fmt.Println(f, g, h, i, j)

}
