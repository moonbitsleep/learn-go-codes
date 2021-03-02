package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes!我在学习了！" // UTF-8
	fmt.Println(s, len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	// 获得字符数量
	fmt.Println("\nRune count:",utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 现在用rune
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()
}
