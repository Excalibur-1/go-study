package main

import "fmt"

func main() {
	// 通过 make 来创建
	//dict := make(map[string]int)
	// 通过字面值创建
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict)

	// 给 map 赋值就是指定合法类型的键，然后把值赋给键
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	colors["Blue"] = "#0000FF"

	// 不初始化 map , 就会创建一个 nil map。nil map 不能用来存放键值对，否则会报运行时错误
	//var colors map[string]string
	//colors["Red"] = "#da1337"
	// Runtime Error:
	// panic: runtime error: assignment to entry in nil map

	//选择是只返回值，然后判断是否是零值来确定键是否存在。
	value := colors["Blue"]
	if value != "" {
		fmt.Println(value)
	}
}
