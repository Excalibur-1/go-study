package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go"}

	value, ok := m[1]
	if ok {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}
}
