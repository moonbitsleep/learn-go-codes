package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent1(nums []int) []int {
	total := make([]int, 101)
	for _, num := range nums { // 统计频率
		total[num]++
	}
	// 从小到大，依次累加,保存的值是小于等于的值
	for i := 1; i < 101; i++ {
		total[i] += total[i - 1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[i] = total[nums[i] - 1]
		}
	}
	return nums
}

func smallerNumbersThanCurrent(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	table := make(map[int]int)


	for i, n := range sorted {
		if _, ok := table[n]; !ok {
			table[n] = i
		}
	}

	sorted = []int {}
	for _, n :=range nums {
		sorted = append(sorted, table[n])
	}
	return sorted
}

func main() {
	nums := []int {8, 1, 2, 2, 3}
	result := smallerNumbersThanCurrent(nums)
	fmt.Println(result)
}