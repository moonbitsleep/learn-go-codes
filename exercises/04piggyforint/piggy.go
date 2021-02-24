package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggybank := 0
	for piggybank < 2000 {
		num := rand.Intn(3) + 1
		switch num {
		case 1:
			piggybank += 5
		case 2:
			piggybank += 10
		case 3:
			piggybank += 25
		}
	}
	fmt.Println(piggybank)
	left := piggybank / 100
	right := piggybank % 100
	fmt.Printf("piggybank:$%2d.%02d\n", left, right)
}
