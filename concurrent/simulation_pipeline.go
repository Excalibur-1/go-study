//模拟管道实现
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")

	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	err := cmd1.Start()
	if err != nil {
		fmt.Println("1 start err:", err)
		return
	}
	err = cmd1.Wait() //开始阻塞
	if err != nil {
		fmt.Println("1 wait err:", err)
		return
	}

	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	err = cmd2.Start()
	if err != nil {
		fmt.Println("2 start err:", err)
		return
	}
	err = cmd2.Wait() //开始阻塞
	if err != nil {
		fmt.Println("2 wait err:", err)
		return
	}

	fmt.Println(outputBuf2.Bytes())
}
