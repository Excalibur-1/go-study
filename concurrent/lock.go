//Go程序可以使用通道进行多个goroutine间的数据交换，但是这仅仅是数据同步中的一种方法。
// Go语言与其他语言如C、Java一样，也提供了同步机制，在某些轻量级的场合，原子访问（sync/atomic包），
// 互斥锁（sync.Mutex）以及等待组（sync.WaitGroup）能最大程度满足需求。
//
//贴士：利用通道优雅的实现了并发通信，但是其内部的实现依然使用了各种锁，因此优雅代码的代价是性能的损失。
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//互斥锁 sync.Mutex
	//互斥锁是传统并发程序进行共享资源访问控制的主要方法。
	//Go中由结构体sync.Mutex表示互斥锁，保证同时只有一个 goroutine 可以访问共享资源。
	//示例1：普通数据加锁
	var mutex sync.Mutex
	num := 0
	//开启10个协程，每个协程都让共享数据num+1
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock() //加锁，阻塞其他协程获取锁
			num += 1
			mutex.Unlock() //解锁
		}()
	}
	//模拟子协程处理结束，等待1秒
	time.Sleep(time.Second)
	// 输出1000，如果没有加锁，则输出的数据很大可能不是1000
	fmt.Println("num=", num)

	/*
		一旦发生加锁，如果另外一个 goroutine 尝试继续加锁时将会发生阻塞，直到这个 goroutine 被解锁，所以在使用互斥锁时应该注意一些常见情况：
		对同一个互斥量的锁定和解锁应该成对出现，对一个已经锁定的互斥量进行重复锁定，会造成goroutine阻塞，直到解锁
		对未加锁的互斥锁解锁，会引发运行时崩溃，1.8版本之前可以使用defer可以有效避免该情况，但是重复解锁容易引起goroutine永久阻塞，
		1.8版本之后无法利用defer+recover恢复
	*/

	//示例2：对象加锁
	a := &Account{
		money: 0,
		lock:  &sync.Mutex{},
	}
	for i := 0; i < 100; i++ {
		go func(num int) {
			a.Add(num)
		}(10)
	}
	//模拟子协程处理结束，等待1秒
	time.Sleep(time.Second)
	a.Query() //不加锁会打印不到1000的数值，加锁后打印1000

	/*
		读写锁 sync.RWMutex
		在开发场景中，经常遇到多处并发读取，一次并发写入的情况，Go为了方便这些操作，在互斥锁基础上，提供了读写锁操作。
		读写锁即针对读写操作的互斥锁，简单来说，就是将数据设定为 写模式（只写）或者读模式（只读）。
		使用读写锁可以分别针对读操作和写操作进行锁定和解锁操作。
		读写锁的访问控制规则与互斥锁有所不同：
			写操作与读操作之间也是互斥的
			读写锁控制下的多个写操作之间是互斥的，即一路写
			多个读操作之间不存在互斥关系，即多路读
		在Go中，读写锁由结构体sync.RWMutex表示，包含两对方法：
		// 设定为写模式：与互斥锁使用方式一致，一路只写
		func (*RWMutex) Lock()				// 锁定写
		func (*RWMutex) Unlock()			// 解锁写

		// 设定为读模式：对读执行加锁解锁，即多路只读
		func (*RWMutex) RLock()
		func (*RWMutex) RUnlock()
	*/
	/*
		注意：
		Mutex和RWMutex都不关联goroutine，但RWMutex显然更适用于读多写少的场景。
		仅针对读的性能来说，RWMutex要高于Mutex，因为rwmutex的多个读可以并存。
		所有被读锁定的goroutine会在写解锁时唤醒
		读解锁只会在没有任何读锁定时，唤醒一个要进行写锁定而被阻塞的goroutine
		对未被锁定的读写锁进行写解锁或读解锁，都会引发运行时崩溃
		对同一个读写锁来说，读锁定可以有多个，所以需要进行等量的读解锁，才能让某一个写锁获得机会，否则该goroutine一直处于阻塞，
		但是sync.RWMutext没有提供获取读锁数量方法，这里需要使用defer避免，如下案例所示。
	*/

	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("Try Lock reading i:", i)
			rwm.RLock()
			fmt.Println("Ready Lock reading i:", i)
			time.Sleep(time.Second)
			fmt.Println("Try Unlock reading i:", i)
			rwm.RUnlock()
			fmt.Println("Ready Unlock reading i:", i)
		}(i)
	}
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Try Lock writing")
	rwm.Lock()
	fmt.Println("Ready Locked writing")
	//上述案例中，只有循环结束，才会执行写锁，所以输出如下：
	//Ready Locked writing		// 总在最后一行

	//读写锁补充 RLocker方法
	//sync.RWMutex类型还有一个指针方法RLocker：
	//func (rw *RWMutex) RLocker() Locker
	//返回值Locker是实现了接口sync.Lokcer的值，该接口同样被 *sync.Mutex和*sync.RWMutex实现，包含方法：Lock和Unlock。
	//当调用读写锁的RLocker方法后，获得的结果是读写锁本身，该结果可以调用Lock和Unlock方法，和RLock，RUnlock使用一致。
	//读写锁的内部其实使用了互斥锁来实现，他们都使用了同步机制：信号量。

	//死锁
	//常见会出现死锁的场景：
	//两个协程互相要求对方先操作，如：AB相互要求对方先发红包，然后自己再发
	//读写双方相互要求对方先执行，自己后执行
	//模拟死锁:程序无法正常输出读取和写入的值
	var rwm2 sync.RWMutex
	ch := make(chan int)
	go func() {
		rwm2.RLock() //加读锁
		x := <-ch    //如果不写入，则无法读取
		fmt.Println("读取到的x:", x)
		rwm2.RUnlock()
	}()
	go func() {
		rwm2.Lock() //加写锁
		ch <- 10    //管道无缓存，没人读走，则无法写入
		fmt.Println("写入：", 10)
		rwm2.Unlock()
	}()
	time.Sleep(time.Second)
	//将上述死锁案例中的锁部分代码去除，则两个协程正常执行。
}

type Account struct {
	money int
	lock  *sync.Mutex
}

func (a *Account) Query() {
	fmt.Println("当前金额为：", a.money)
}

func (a *Account) Add(num int) {
	a.lock.Lock()
	a.money += num
	a.lock.Unlock()
}
