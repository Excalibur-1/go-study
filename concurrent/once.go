//Once 只执行一次
//sync包除了提供了互斥锁、读写锁、条件变量外，
//还提供了一个结构体：sync.Once，负责只执行一次，也即全局唯一操作。
package main

import (
	"fmt"
	"sync"
)

func main() {
	//使用方式如下：
	var once sync.Once
	once.Do(func() {}) // Do方法的有效调用次数永远是1

	//sync.Once的典型应用场景是只执行一次的任务，如果这样的任务不适合在init函数中执行，该结构体类就会派上用场。
	//sync.Once内部使用了“卫述语句、双重检查锁定、共享标记的原子操作”来实现Once功能。

	//once示例：
	var once1 sync.Once
	var wg sync.WaitGroup

	//初始化信息
	p := &Person{
		Name: "比尔",
		Age:  0,
	}

	//启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once1.Do(func() {
				p.Grown()
			})
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("年龄是：", p.Age) //只会输出1
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grown() {
	p.Age += 1
}
