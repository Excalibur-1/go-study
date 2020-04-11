//定时器
package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器time.Timer的底层其实也是一个管道。
	//time.Timer
	//在时间到达之前，没有数据写入，则timer.C会一直阻塞，直到时间到达，系统会自动向timer.C中写入当前时间，阻塞就会被解除：
	//创建定时器，定义延迟时间为2秒
	layTimer := time.NewTimer(time.Second * 2)
	//从管道取数据，但是一致都是空的，阻塞中，直到2秒后有数据才能取出
	fmt.Println(<-layTimer.C)

	//定时器的其他操作：
	//Stop()：停止定时器，此时如果从管道中取数据，则会阻塞
	//Reset()：重置定时器，此时需要传入一个新的定时时间间隔
	//timer.Tiker：可以创建一个周期定时器
}
