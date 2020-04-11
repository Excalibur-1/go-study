//等待组 sync.WaitGroup
//sync.WaitGroup类型的值也是并发安全的，该类型结构体中内内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少 。
//当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减1。
//同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时， 表示所有任务己经完成。
//等待组常用方法：
//(wg *WaitGroup) Add(delta int) 等待组计数器+1，该方法也可以传入负值让等待计数减少，切记不能减少为负数，会引发崩溃
//(wg *WaitGroup) Done() 等待组计数器-1，等同于Add传入负值
//(wg *WaitGroup) Wait() 等待组计数器!=0时阻塞，直到为0
//应用场景：WaitGroup一般用于协调多个goroutine运行。
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup //声明等待组
	var urls = []string{
		"https://www.baidu.com",
		"https://www.163.com",
		"https://www.weibo.com",
	}
	for _, url := range urls {
		wg.Add(1) //每个任务开始，等待组+1
		go func(url string) {
			defer wg.Done()         //任务执行完成，等待组-1
			_, err := http.Get(url) //执行访问
			if err != nil {
				fmt.Println("http.Get err:", err)
			}
			fmt.Println(url)
		}(url)
	}
	wg.Wait() //等待所有任务完成
	fmt.Println("over")
	//上述案例可以使用channel方式，每个go协程执行时，channel传递完成信号，但是使用通道的方式明显过重。
	//贴士：有了等待组，我们就不需要再在main函数中使用 time.Sleep()方法来模拟等待协程运行结束了。

	//等待组与锁配合使用示例：
	//开启多个协程对共享内存产生竞争，单独使用等待组不能解决问题，如下所示：
	var wg1 sync.WaitGroup
	var money = 10000
	//开启10个协程。每个协程内部循环100次，每次循环+10
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				money += 10 //多个协程对money产生了竞争
			}
			wg1.Done()
		}()
	}
	wg1.Wait()
	fmt.Println("最终的money=", money) // 应该输出20000才正确，但是每次运行输出结果不正确

	//等待组与互斥锁（同步锁）配合解决钱数问题示例：
	//开启多个协程对共享内存产生竞争，单独使用等待组不能解决问题，如下所示：
	var mt sync.Mutex
	var wg2 sync.WaitGroup
	var money2 = 10000
	//开启10个协程。每个协程内部循环100次，每次循环+10
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(index int) {
			mt.Lock()
			//fmt.Printf("协程%d抢到锁\n",index)
			for j := 0; j < 100; j++ {
				money2 += 10 //多个协程对money产生了竞争
			}
			//fmt.Printf("协程%d准备解锁\n",index)
			mt.Unlock()
			wg2.Done()
		}(i)
	}
	wg2.Wait()
	fmt.Println("最终的money=", money2) // 应该输出20000才正确

}
