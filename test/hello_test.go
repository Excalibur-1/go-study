package test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"study/test/hello"
	"testing"
)

//代码覆盖率命令：
//go test -v -cover

//go test -v test/hello_test.go 							# -v用于显示详细测试流程
//go test -v -run Test_hello test/hello_test.go 			# 只执行Test_hello
func TestHello(t *testing.T) {
	r := hello.Hello()
	if r != "world" {
		t.FailNow()
	}
}

//断言库
//使用一些第三方的断言库也可以达到原生的单元测试效果：
func TestHello2(t *testing.T) {
	r := hello.Hello()
	assert.Equal(t, "world", r)
}

//BDD测试框架：goConvey
func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	convey.Convey("Given some integer with a starting value", t, func() {
		x := 1

		convey.Convey("When the integer is incremented", func() {
			x++

			convey.Convey("The value should be greater by one", func() {
				convey.So(x, convey.ShouldEqual, 2)
			})
		})
	})
}

//基准测试：用于测试一段程序的运行性能及CPU消耗
//性能测试函数以 Benchmark 为名称前缀，同样保存在 *_test.go 文件里。示例：
func Benchmark_Hello(b *testing.B) {
	//开始测试性能相关代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
	}

	//结束性能测试
	b.StopTimer()
	//go test -v -bench=Hello  		# -bench=.表示运行所有基准测试。win下参数为：-bench="-"
	//-benchmem 			# 显示性能具体的开销情况
	//-benchtime=5s		# 自定义测试时间为5秒
}
