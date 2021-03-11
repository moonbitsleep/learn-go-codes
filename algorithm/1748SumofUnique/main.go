package main

import "fmt"

func sumOfUnique(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	sum := 0
	for k, v := range m {
		if v == 1 {
			sum += k
		}
	}
	return sum
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 7, 9}
	fmt.Println(sumOfUnique(nums))
}
