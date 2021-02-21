package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
func split10(sum int) (ten, one int) {
	ten = sum / 10
	one = sum % 10
	return
}

func main() {
	fmt.Println("3 + 5 =", add(3, 5))
	a := "kick"
	b := "ass"
	a, b = swap(a, b)
	fmt.Println(a, b)
	x, y := split10(29)
	fmt.Println("The split of x = ", x, "The split of y = ", y)
}
