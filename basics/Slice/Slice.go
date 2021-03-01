package main

import "fmt"
/***
Slice 本身没有数据，是对底层array的一个view
 */
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2)

	fmt.Println("After update Slice")
	updateSlice(s1)
	// 切片发生改变，原数组跟着发生改变
	fmt.Println("arr[2:] = ", s1)
	fmt.Println("arr = ", arr)

	fmt.Println("Reslice...")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//切片的扩展，仅能向后不可向前
	// slice的数据结构：ptr len cap，扩展不能超过cap(s),取值不能超过len(s)
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v \n",s1, len(s1), cap(s1)) // 2,3,4,5
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v \n",s2, len(s2), cap(s2)) // 5,6

}
