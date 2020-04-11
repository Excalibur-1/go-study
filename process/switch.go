package main

import "fmt"

func main() {
	result := 5
	switch {
	case (result >= 0 && result < 5):
		fmt.Println("result > 0")
	case (result >= 5 && result < 10):
		fmt.Println("result > 5")
	case (result >= 10 && result < 15):
		fmt.Println("result > 10")
	case (result >= 15 && result < 20):
		fmt.Println("result > 15")
	default:
		fmt.Println("error result")
	}
}
