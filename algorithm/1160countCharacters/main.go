package main

import (
	"fmt"
	"strings"
)

func countCharacters1(words []string, chars string) int {
	// chars中的只能用一次
	var wholeLength int = 0
	OutLoop: for _, v1 := range words {
	// 拿每一个单词。
	for _, v2 := range v1{
		// 取字母
		if strings.Count(v1, string(v2)) > strings.Count(chars, string(v2)) {
			// 如果在单词中某个字母的数量大于词库中的数量，说明弄不了。
			continue OutLoop
		}
	}
	// 如果这个循环走完了，能走到这里，说明这个单词里的字母都在chars中
	wholeLength += len(v1)
}
	return wholeLength
}

// 使用map解决的方法
func countCharacters(words []string, chars string) int {
	var res int
	m := make(map[byte]int)

	// 对字母表的字符进行统计
	for i :=range chars {
		// golang string的元素是byte(uint8)
		m[chars[i]]++
	}
	for _, w :=range words {
		if isContainer(w,m) {
			res += len(w)
		}
	}
	return res
}

func isContainer(w string, m map[byte]int) bool {
	charMap := map[byte]int{}
	// 统计一个单词中的字符个数
	for i :=range w {
		charMap[w[i]]++
	}

	// 对char_map 和 m（字符表）进行比较，判断是否具有足够对字母
	for k, v :=range charMap {
		if m[k] < v {
			return false
		}
	}
	return true
}

func main() {
	words := []string {"hello", "world", "leetcode"}
	chars := "welldonehoneyr"
	// 实验代码
	//for i :=range chars {
	//	fmt.Printf("%[1]T : %[1]q\t", chars[i])
	//}
	//var c byte = 'w'
	//fmt.Printf("\n c = {%[1]T : %[1]q}\n", c)
	//
	res := countCharacters(words, chars)
	fmt.Println(res)
}
