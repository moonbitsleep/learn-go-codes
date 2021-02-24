package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggybank float64 = 0.0
	count5, count10, count25 := 0, 0, 0
	for piggybank < 20.0 {
		num := rand.Intn(3) + 1
		switch num {
		case 1:
			piggybank += 0.05
			count5++
		case 2:
			piggybank += 0.10
			count10++
		case 3:
			piggybank += 0.25
			count25++
		}

		fmt.Printf("存钱罐有:%7.2f美元\n", piggybank)
		fmt.Printf("有%d枚5美分硬币\n", count5)
		fmt.Printf("有%d枚10美分硬币\n", count10)
		fmt.Printf("有%d枚25美分硬币\n", count25)
		fmt.Println("------------------------")
	}
}
