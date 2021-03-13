package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("The time is ", time.Now(), " Welcome to go playground!")
	fmt.Println("My Number is ", rand.Intn(16))
}
