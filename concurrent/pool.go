//对象池 sync.Pool
//sync.Pool类型可以视为临时值的容器，该容器具备自动伸缩、高效特性，同时也是并发安全的，其方法有：
//Get：从池中取出一个interface{}类型的值
//Put：把一个interface{}类型的值存于池中
//注意：
//如果池子从未Put过，其New字段也没有被赋值一个非nil值，那么Get方法返回结果一定是nil。
//Get获取的值不一定存在于池中，如果Get到的值存在于池中，则该值Get后会被删除
//对象池原理：
//对象池可以把内部的值产生的压力分摊，即专门为每个操作它的协程关联的P建立本地池。
//Get方法被调用时，先从本地P对象的本地私有池和本地共享池获取值，如果获取失败，则从其他P的本地私有池偷取一个值返回，
//如果依然没有获取到，会依赖于当前对象池的值生成函数。注意：生产函数产生的值不会被放到对象池中，只是起到返回值作用
//对象池的Put方法会把参数值存放到本地P的本地池中，每个P的本地共享池中的值，都是共享的，即随时可能会被偷走
//对象池对垃圾回收友好，执行垃圾回收时，会将对象池中的对象值全部溢出
//应用场景：存储被分配了但是未被使用，未来可能会被使用的值，以减少GC的压力。
//案例：由于fmt包总是使用一些[]byte对象，golang为期建立了一个临时对象池，存放这些对象，
//需要的时候，从pool中取，拿不到则分配一份，这样就能避免一直生成[]byte，垃圾回收的效率也高了很多。
package main

import (
	"fmt"
	"sync"
	"time"
)

//示例：
//一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()
	//不使用对象池
	for i := 0; i < 100000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	//使用对象池
	for i := 0; i < 100000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool", b-a, "s")
	fmt.Println("with pool", c-b, "s")
}
