package main

import (
	"bufio"
	"fmt"
	"os"
)

// 栈结构
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("123")
	//fmt.Println(4)
}

// defer从下向上
func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()


	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, i)
	}
}

// 参数在defer语句时计算
func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("err!")
		}
	}
}

func main() {
	//tryDefer()
	//writeFile("1-20.txt")
	tryDefer2()
}
