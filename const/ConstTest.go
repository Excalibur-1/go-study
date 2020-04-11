package main

import (
	"fmt"
	"math"
)

//常量定义
const A = 3
const PI float32 = 3.1415
const MASK = 1 << 3

const (
	deadbeef = 0xdeadbeef        // untyped int with value 3735928559
	a        = uint32(deadbeef)  // uint32 with value 3735928559
	b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	c        = float64(deadbeef) // float64 with value 3735928559 (exact)
	//d = int32(deadbeef)   // compile error: constant overflows int32
	//e = float64(1e309)    // compile error: constant overflows float64
	//f = uint(-1)          // compile error: constant underflows uint
)

//无类型常量， math.Pi无类型浮点数常量，可以直接用于任意需要浮点数或复数的地方
var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi

//指定math.Pi为特定类型
const Pi64 float64 = math.Pi

//则对需要使用float32和complex128的地方就需要明确的类型转换
var a1 float32 = float32(Pi64)
var b1 float64 = Pi64
var c1 complex128 = complex128(Pi64)

func main() {
	//除法运算符会根据操作数的类型生成对应类型的结果
	var f float64 = 212
	// "100"; (f - 32) * 5 is a float64
	fmt.Println((f - 32) * 5 / 9)
	// "0";   5/9 is an untyped integer, 0
	fmt.Println(5 / 9 * (f - 32))
	// "100"; 5.0/9.0 is an untyped float
	fmt.Println(5.0 / 9.0 * (f - 32))

	//无类型常量隐式转换为等号左边的类型
	var f1 float64 = 3 + 0i // untyped complex -> float64
	f1 = 2                  // untyped integer -> float64
	f1 = 1e123              // untyped floating-point -> float64
	f1 = 'a'                // untyped rune -> float64

	//上面的语句相当于
	/*var f1 float64 = float64(3 + 0i)
	f1 = float64(2)
	f1 = float64(1e123)
	f1 = float64('a')*/
	fmt.Println(f1)

	//没有显示声明类型的变量声明，常量的形式决定变量的默认类型
	i := 0      // untyped integer;        implicit int(0)
	r := '\000' // untyped rune;           implicit rune('\000')
	f2 := 0.0   // untyped floating-point; implicit float64(0.0)
	c := 0i     // untyped complex;        implicit complex128(0i)
	fmt.Println(i, r, f2, c)

	//给变量设置不同类型
	/*var i = int8(0)
	var i int8 = 0*/

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}
