package main

import (
	"fmt"
	"math/cmplx"
)

// 全局变量
var c, java bool

// := 结构不能在函数外使用
//dd := 123

func convert() {
	var i int64 = -105
	var j float64 = float64(i)
	var u uint = uint(j)
	fmt.Println(i, j, u)
}

// 未指定类型的常量由上下文来决定其类型
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var python bool
	var i int32
	var f float32
	var c rune
	// 0 false false 0 0 0
	fmt.Println(c, java, python, i, f, c)

	var implicit = '天'
	fmt.Println(implicit)        // 109
	fmt.Printf("%c\n", implicit) // 天
	var java = "java is complex"
	fmt.Println(java)
	// int, uint 和 uintptr 位宽与机器位宽相等
	// byte // uint8 的别名
	// rune // int32 的别名 表示一个 Unicode 码点
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 零值
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 类型转换
	fmt.Println("类型转换")
	convert()

	// 常量不能用 := 语法声明
	//const hello := "nice"
	const world = "世界"

	//Big 和 Small是未指定类型(untyped)的数值常量
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
