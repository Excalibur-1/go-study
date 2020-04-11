package package_test

import "fmt"

func Hello() {
	fmt.Println("Hello from Go.")
	hi()
}

func hi() {
	fmt.Println("hi from Go.")
}
