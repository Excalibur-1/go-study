package main

import "fmt"

func main() {
	var b bool
	b = true
	fmt.Printf("b is of type %t\n", b)
	e := bool(true)
	fmt.Printf("e is of type %t\n", e)
}
