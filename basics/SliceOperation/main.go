package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func main() {
	arr := [...]int {0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// s4, s5已经不再是arr的view
	fmt.Println("arr = ", arr)

	var s []int //零值为nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println("s = ", s)
	s1 = []int{2, 4, 6, 8}
	printSlice(s1)

	s2 = make([]int, 16)
	printSlice(s2)
	s3 = make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Copying...")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("deleting.....")
	// 删除中间元素 8
	//...代表可变参数
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("popping from front....")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front = ", front)
	printSlice(s2)

	fmt.Println("popping from back....")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println("tail = ", tail)
	printSlice(s2)
}
