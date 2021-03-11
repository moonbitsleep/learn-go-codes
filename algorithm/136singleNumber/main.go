package main

import "fmt"

func singleNumber1(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	fmt.Println(m)
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 0 xor 甲 = 甲, 甲 xor 甲 = 0， 根据序列特点，唯一值xor到最后会剩下
func singleNumber(nums []int) int {
	xor := 0

	for _, v := range nums {
		xor ^= v
	}
	return xor
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
