package main

import "fmt"

// Vertex for a new struct Type
type Vertex struct {
	X int
	Y int
}

// 与 C 不同，Go 没有指针运算。
func pointer() {
	// 类型 *T 是指向 T 类型值的指针。其零值为 nil。
	var p *int
	fmt.Println("p的初值为", p)
	// & 操作符会生成一个指向其操作数的指针。
	i, j := 42, 2701

	p = &i                         // 指向 i
	fmt.Println("通过指针读取 i 的值", *p) // 通过指针读取 i 的值
	*p = 21                        // 通过指针设置 i 的值
	fmt.Println("查看 i 的值", i)      // 查看 i 的值

	p = &j                    // 指向 j
	*p = *p / 37              // 通过指针对 j 进行除法运算
	fmt.Println("查看 j 的值", j) // 查看 j 的值
}

func structType() {
	v := Vertex{1, 2}
	fmt.Println(v)
	// 结构体字段使用点号来访问
	v.X = 4
	fmt.Println(v.X)
	// 结构体字段可以通过结构体指针来访问
	ptr := &v
	// 使用隐式间接引用
	ptr.Y = 9
	// 显示间接引用，过于复杂
	(*ptr).X = 8
	fmt.Println("ptr:", ptr, "*ptr:", *ptr)
}

func structSignValue() {
	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)
	fmt.Println("创建结构体:", v1, p, v2, v3)
}

func arrayDisplay() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "golang"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:3]
	fmt.Println(s)

	// 更改切片的元素会修改其底层数组中对应的元素
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// 切片文法类似于没有长度的数组文法。
func slice2() {
	q := []int{2, 3, 5, 7, 11, 13, 17}
	r := []bool{true, false, false, true}
	s := []struct {
		i int64
		j float64
	}{
		{1, 1.1},
		{2, 1.2},
		{3, 1.3},
	}
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
}

// 切片的长度就是它所包含的元素个数。
// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
func sliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)

	//切片的零值是 nil
	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("nil!")
	}
}

// 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本
func forRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// 可以将下标或值赋予 _ 来忽略它。
//只需要索引，忽略第二个变量即可。 for i := range pow
func rangeIgnore() {
	var pow = make([]int, 10)
	//只需要索引，忽略第二个变量
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// 忽略第一个变量
	for _, val := range pow {
		fmt.Printf("%d\n", val)
	}
}

func mapping() {
	type Vertex struct {
		Lat, Long float64
	}
	// 映射的零值为 nil 。nil 映射既没有键，也不能添加键。
	var m map[string]Vertex
	// make 函数会返回给定类型的映射，并将其初始化备用。
	m = make(map[string]Vertex)
	// 映射将键映射到值。
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}

func mappingDef() {
	type Vertex struct {
		Lat, Long float64
	}
	var m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mappingCrud() {
	m := make(map[string]int)
	// 在映射 m 中插入或修改元素,类似Python字典
	m["Answer"] = 19
	fmt.Println("The value of m[\"Answer\"]", m["Answer"])
	//获取元素
	m["Answer"] = 48
	elem := m["Answer"]
	fmt.Println("The value of m[\"Answer\"]", elem)
	// 删除元素：delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	/*
		通过双赋值检测某个键是否存在：
			elem, ok := m[key]
		若 key 在 m 中，ok 为 true ；否则，ok 为 false。

		若 key 不在映射中，那么 elem 是该映射元素类型的零值。

		同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
	*/
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	pointer()
	structType()
	structSignValue()
	arrayDisplay()
	slice()
	slice2()
	sliceLenCap()
	forRange()
	rangeIgnore()
	mapping()
	mappingDef()
	mappingCrud()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
