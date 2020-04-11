package main

import (
	"fmt"
)

func main() {
	a := 10
	var p *int = &a
	fmt.Println(*p)
	a++
	fmt.Println(a)

	if 1 < 2 {
		fmt.Println("ok")
	}

	if a := 1; a < 2 {
		fmt.Println(a, "ok")
	}

	fmt.Println(a, "ok")

	b := 1
	for {
		if b > 3 {
			break
		}
		fmt.Println(b)
		b++
	}
	fmt.Println("over")

	b = 1
	for b <= 3 {
		fmt.Println(b)
		b++
	}
	fmt.Println("over")

	b = 1
	for i := 0; i < 3; i++ {
		fmt.Println(b)
		b++
	}
	fmt.Println("over")

	b = 1
	switch b {
	case 0:
		fmt.Println("b=0")
	case 1:
		fmt.Println("b=1")
	default:
		fmt.Println("none")
	}
	fmt.Println("over")

	switch {
	case b >= 0:
		fmt.Println("b=0")
	case b >= 1:
		fmt.Println("b=1")
	default:
		fmt.Println("none")
	}
	fmt.Println("over")

	switch {
	case b >= 0:
		fmt.Println("b=0")
		fallthrough
	case b >= 1:
		fmt.Println("b=1")
	default:
		fmt.Println("none")
	}
	fmt.Println("over")

	switch c := 1; {
	case c >= 0:
		fmt.Println("c=0")
		fallthrough
	case c >= 1:
		fmt.Println("c=1")
	default:
		fmt.Println("none")
	}
	fmt.Println("over")

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//break指定标签名跳出，此处如果不指定标签名则会无限循环，break只能跳出一层循环
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//使用goto关键字跳出循环，必须将标签名指定在循环语句之后，这样才能避免死循环
				goto LABEL2
			}
		}
	}
LABEL2:
	fmt.Println("OK")

LABEL3:
	for i := 0; i < 10; i++ {
		for {
			//continue配合标签直接跳转到外层的有限循环执行，虽然内部是无线循环但也不会死循环
			continue LABEL3
			fmt.Println("NO")
		}
	}
	fmt.Println("OK")

}
