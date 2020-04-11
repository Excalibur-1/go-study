package main

import "fmt"

//常量定义
const a int = 1
const b = 2
const c, d, e, f = 1, "1", true, 1.1
const (
	g, h, i = 2, "4", 5.5
)
const (
	j = 6
	k
	l
	m    = len(h)
	n, o = 1, "2"
	p, q
)

const (
	r = 'A'
	s
	t = iota
	u
)

const (
	v = iota
)

const (
	w float64 = 1 << (iota ^ 10)
	x
	y
	z
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m)
	fmt.Println(n, o, p, q)
	fmt.Println(r, s, t, u, v)
	fmt.Println(w, x, y, z)
}
