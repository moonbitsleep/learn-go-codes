package main

import (
	"fmt"
	"strings"
)

func countWords(article string) map[string]int {
	article = strings.ToLower(article)
	wordsWithOp := strings.Fields(article)

	// 创建一个列表保存不带标点的元素
	var words []string
	for _, word :=range wordsWithOp {
		words = append(words, strings.Trim(word, ",.!;?"))
	}

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}


func countPrint(f map[string]int) {
	for t, num := range f {
		fmt.Printf(" [%v] occurs %d times \n", t, num)
	}
}

func main() {
	article := "He ate the food and wine in the Stomach! would turn; into Bile, blood clotting? into the ice, the heart as hard as iron."
	//article = strings.ToLower(article)
	//words := strings.Fields(article)
	//var wordsWithoutOp []string
	//for _, word := range words {
	//	wordsWithoutOp = append(wordsWithoutOp, strings.Trim(word, ",.!;?"))
	//}
	//fmt.Printf("%q\n", wordsWithoutOp)
	//
	//frequency := make(map[string]int)
	//
	//for _, word := range wordsWithoutOp {
	//	frequency[word]++
	//}
	//
	//for t, num := range frequency {
	//	fmt.Printf(" [%v] occurs %d times \n", t, num)
	//}

	fre := countWords(article)
	countPrint(fre)
}
