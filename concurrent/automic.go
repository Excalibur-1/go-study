//原子操作
//对并发的操作，可以利用Go新提出的管道思想，大多数语言都有的锁思想，也可以使用最基础的原子操作。
//原子操作的执行过程不能被中断，因为此时CPU不会去执行其他对该值进行的操作。
//Go语言提供的原子操作是非侵入式的，由标准库代码包sync/atomic提供了一系列原子操作函数。
//这些函数可以对一些数据类型进行原子操作：int32，int64，uint32，uint64，uintptr，unsafe.Pointer。
//这些函数提供的原子操作有5种：增、减、比较并交换、载入、存储、交换。
//注意：
//不能取地址值的数值无法进行原子操作
//在性能上，原子操作比互斥锁性能更高
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//常用原子操作
	//原子运算：增/减
	var i32 int32
	i32 = 10
	//增加函数的函数名前缀都是Add开头
	// 原子性的把一个int32类型变量 i32 增大3 ，下列函数返回值必定是已经被修改的值
	newi32 := atomic.AddInt32(&i32, 3) // 传入指针类型因为该函数需要获得数据的内存位置，以施加特殊的CPU指令
	fmt.Println(newi32)
	/*
		常见的增/减原子操作函数：
			atomic.AddInt32
			atomic.AddInt64
			atomic.AddUint32
			atomic.AddUint64
			atomic.AddUintptr
		注意:
		如果需要执行减少操作，可以这样书写 atomic.AddInt32(&i32, -3)
		对uint32执行增加NN（代表负整数，增加NN也可以理解为减少-NN）：
			atomic.AddUint32(&ui32, ^uint32(-NN-1)) //原理：补码
		对uint64以此类推
		不存在atomic.AddPointer的函数，因为unsafe.Poniter类型的值无法被增减
	*/

	//原子运算：比较与替换
	//比较并替换即“Compare And Swap”，简称CAS。该类原子操作名称都以CompareAndSwap为前缀。
	//举例：
	//参数一：被操作数 参数二和参数三代表被操作数的旧值和新值
	//func CompareAndSwapInt32(addr *int32, old, new int32)(swap bool)

	//CAS的一些特点：
	//CAS与锁相比，明显不同是它总是假设操作值未被改变，一旦确认这个假设为真，立即进行替换。所以锁的做法趋于悲观，CAS的做法趋于乐观。
	//CAS的优势：可以在不创建互斥量和不形成临界区的情况下，完成并发安全的值替换操作，可以大大减少性能损耗。
	//CAS的劣势：在操作频繁变更的情况下，CAS操作容易失败，有时候需要for循环判断返回结构的bool来进行多次尝试
	//CAS操作不会阻塞协程，但是仍可能使流程执行暂时停滞（这种停滞极短）
	//应用场景：并发安全的更新一些类型的值，可以优先选择CAS操作。

	//原子读取：载入
	//为了原子的读取数值，Go提供了一系列载入函数，名称以Load为前缀。
	//CAS与载入的配合示例：
	var value int32 = 20
	var num int32 = 10
	AddValue(value, num)

	/*
		原子写入：存储
		在原子存储时，任何CPU都不会进行针对同一个值的读写操作，此时不会出现：在并发时，别人读取到了修改了一半的值。
		Go的sync/atomic包提供的存储函数都是以Store为前缀。
		示例：
		// 参数一位被操作数据的指针 参数二是要存储的新值
		atomic.StoreInt32(i *int3, v int32)
		Go原子存储的特点：存储操作总会成功，因为不关心被操作值的旧值是什么，这与CAS有明显区别。
	*/

	/*
		交换
		交换与CAS操作相似，但是交换不关心被操作数据的旧值，而是直接设置新值，不过会返回被操作值的旧值。交换操作比CAS操作的约束更少，且比载入操作功能更强。
		在Go中，交换操作都以Swap为前缀，示例：
		// 参数一是被操作值指针  参数二是新值  返回值为旧值
		atomic.SwapInt32(i *int32, v int32)
	*/

	/*
		原子值
		Go还提供了sync/atomic.Value类型的结构体，用于存储需要原子读写的值。该类型与第二章中的原子操作最大的区别是，可接受的值类型不限。
		示例：
		    var atomicV atomic.Value
		该结构体包含的方法：
			Load：原子的读取原子值实例的值，返回interface{}类型结果
			Store：原子的向原子值实例中存储值，接受一个interface{}类型参数（不能是nil），无返回结果
		注意：
		如果一个原子值没有通过Store方法存储值，那么其Load方法总是返回nil
		原子值实例一旦存储了一个类型的值，后续再次Store存储时，存储的值必须也是原有的类型
		尤其注意：atomic.Value变量（指针类型变量除外）声明后，值不应该被赋值到别处，
		比如赋值给别的变量，作为参数值传入函数，作为结果值从函数返回等等，这样会有安全隐患。
		因为结构体值的复制不但会生产该值的副本，还会生成其中字段的副本，会造成并发安全保护失效。
	*/

}

func AddValue(value, num int32) {
	for {
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, v+num) {
			break
		}
	}
	fmt.Println(value)
}
