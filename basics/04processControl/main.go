package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sum100() int {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func whileSum100() int {
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	return sum
}

func mysqrt(x float64) string {
	if x < 0 {
		return mysqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句
func pow(x, n, lim float64) float64 {
	// 找一个较小的数字
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Sqrt for me
func Sqrt(x float64) float64 {
	var z = float64(x / 2)
	var count int64 = 1
	for math.Abs((z*z-x)/(2*z)) > 1e-5 {
		count++
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println("Sqrt 循环次数为", count)
	return z
}

func main() {
	// 循环
	fmt.Println("1 + 2 + 3 + ... + 100 = ", sum100())
	fmt.Println("等价while写法", whileSum100())
	// 无限循环
	// for {
	// }
	fmt.Println(mysqrt(2), mysqrt(-4))
	// 9， 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	s1 := Sqrt(29)
	s2 := math.Sqrt(29)
	fmt.Println(s1, s2)

	codeSwitch()
	daysForSaturday()
	timesForNote()

	// defer 语句会将函数推迟到外层函数返回之后执行。
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer fmt.Println("\n推迟的第二行")
	fmt.Println("main中要求打印的第一行")

	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	countingDefer()
}

//  switch 的 case 无需为常量，且取值不必为整数
//	这种形式能将一长串 if-then-else 写得更加清晰。
func codeSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func daysForSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("Today :", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//没有条件的 switch 同 switch true 一样
func timesForNote() {
	t := time.Now()
	println("Now Time:", t.String())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good night!")
	}
}

func countingDefer() {
	fmt.Println("counting defer")

	for i := 1; i <= 10; i++ {
		defer fmt.Printf("%d  ", i)
	}
	fmt.Println("done")
}
