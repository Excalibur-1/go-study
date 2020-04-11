package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := array[:]
	fmt.Println("s = ", s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s1 := array[3:6]
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	s1 = append(s1, 8)
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
}
