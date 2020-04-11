//条件变量
//sync.Cond类型即是Go中的条件变量，该类型内部包含一个锁接口。条件变量通常与锁配合使用：
//创建条件变量的函数：
//func NewCond(l locker) *Cond        // 条件变量必须传入一个锁，二者需要配合使用
//*sync.Cond类型有三个方法：
//Wait: 该方法会阻塞等待条件变量满足条件。也会对锁进行解锁，一旦收到通知则唤醒，并立即锁定该锁
//Signal: 发送通知(单发)，给一个正在等待在该条件变量上的协程发送通知
//Broadcast: 发送通知(广播），给正在等待该条件变量的所有协程发送通知，容易产生 惊群
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	condition := false
	//开启一个新的协程，修改变量condition
	go func() {
		time.Sleep(time.Second)
		cond.L.Lock()    //获取到锁才能改数据
		condition = true //状态变更，发送通知
		cond.Signal()    //发信号
		cond.L.Unlock()
	}()
	//main协程是被通知的对象，等待通知
	cond.L.Lock()
	for !condition {
		cond.Wait() //内部释放了锁（释放后，子协程能拿到锁），并等待通知（消息）
		fmt.Println("获取到了消息")
	}
	cond.L.Unlock() //// 接到通知后，会被再次锁住，所以要在需要的场合释放
	fmt.Println("运行结束")

	//使用条件变量优化生产消费模型（支持多个生产者、多个消费者）：
	rand.Seed(time.Now().UnixNano()) //设置随机数种子
	//生产消费模型中的通道
	ch := make(chan int, BUF_LEN)
	//启动10个生产者
	for i := 0; i < 10; i++ {
		go Producer(ch)
	}
	//启动10个消费者
	for i := 0; i < 10; i++ {
		go Consumer(ch)
	}

	//阻塞主程序退出
	for {

	}
}

//定义缓冲区大小
const BUF_LEN = 5

//定义全局变量
var cond1 *sync.Cond = sync.NewCond(&sync.Mutex{})

//生产者
func Producer(ch chan<- int) {
	for {
		cond1.L.Lock() //全局条件变量加锁
		length := len(ch)
		//缓冲区满，则等待消费者消费，注意这里不能是if
		for length == BUF_LEN {
			cond1.Wait() //挂起当前协程，等待条件变量满足，唤醒消费者
		}
		ch <- rand.Intn(1000) //写入缓冲区一个随机数
		cond1.L.Unlock()      //生产结束，解除互斥锁
		cond1.Signal()        //一旦生产后，就唤醒其他被阻塞的消费者
		time.Sleep(time.Second * 2)
	}
}

//消费者
func Consumer(ch <-chan int) {
	for {
		cond1.L.Lock() //全局条件变量加锁
		length := len(ch)
		//如果缓冲区为空，则等待生产者生产，注意这里不能是if
		for length == 0 {
			cond1.Wait() //挂起当前协程，等待条件变量满足，唤醒生产者
		}
		fmt.Println("Receive:", <-ch)
		cond1.L.Unlock()
		cond1.Signal()
		time.Sleep(time.Second * 1)
	}
}
