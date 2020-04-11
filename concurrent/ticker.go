package main

import "fmt"
import "time"

func main() {
	//创建Ticker，每秒向管道写数据
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		//从管道取数据
		<-ticker.C

		i++
		fmt.Println("i = ", i)

		if i == 5 {
			//停止定时器，跳出循环
			ticker.Stop()
			break
		}
	}
}
