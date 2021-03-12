package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	m := make(map[string]int)
	var res []string

	words := strings.Fields(A + " " + B)
	for _, word := range words {
		m[word]++
	}

	for _, word := range words{
		if m[word] == 1 {
			res = append(res, word)
		}
	}
	return  res
}



func main() {
	fmt.Println(uncommonFromSentences(" ", " "))
}
