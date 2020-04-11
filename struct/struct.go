package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// 结构体字段使用点号来访问。
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// 结构体字段可以通过结构体指针来访问。
	e := Vertex{1, 2}
	p1 := &e
	p1.X = 1e9
	fmt.Println(e)

	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		p  = &Vertex{1, 2} // 类型为 *Vertex , 特殊的前缀 & 返回一个指向结构体的指针
	)
	fmt.Println(v1, v2, v3, p)
}
