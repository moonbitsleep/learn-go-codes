package main

func twoSum(nums []int, target int) []int {
	table := map[int]int{}

	for i, num := range nums {
		margin := target - num
		// 找差值的index
		marginIndex, ok := table[margin]
		if ok {
			return []int{marginIndex, i}
		}
		// 保存nums的每一个暂未匹配到的元素到字典中
		// 实质是在存差值，便于后期寻找，key为元素，value为索引值
		table[num] = i
	}

	return []int{}
}
