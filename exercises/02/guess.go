package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(60)
	var myNum = 78
	var randomNum = rand.Intn(100) + 1
	for randomNum != myNum {
		fmt.Println("the random Number is ", randomNum)
		if randomNum > myNum {
			fmt.Println("Larger!")
		} else {
			fmt.Println("Smaller!")
		}
		randomNum = rand.Intn(100) + 1
	}
	fmt.Println("the random Number is ", randomNum)
	fmt.Println("Equal!")
}
