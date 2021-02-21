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

func main() {
	pointer()
	structType()
	structSignValue()
}
