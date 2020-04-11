package main

func main() {
	//定义channel，双向的
	ch := make(chan int)

	//双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch //只能写，不能读
	var readCh <-chan int = ch  //只能读，不能写

	writeCh <- 666 //写
	// <-writeCh//写管道不能读 invalid operation: <-writeCh (receive from send-only type chan<- int)

	<-readCh
	// readCh <- 666//读管道不能写 invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	//单向管道不能转换为双向管道
	// var ch2 chan int = writeCh// cannot use writeCh (type chan<- int) as type chan int in assignment

}
