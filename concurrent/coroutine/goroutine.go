package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask goroutine")
		time.Sleep(time.Second)
	}
}

func main() {

	go newTask()

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}
}
