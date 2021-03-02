package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

// node的方法
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 传入指针，才能改变值
func (node *treeNode) setValue(v int) {
	node.value = v
}

func createNode(value int) *treeNode {
	// 返回局部变量的地址，不会出错
	//分配到堆上面，自动回收
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	// 因为go语言允许nil传进去，所以不必判断node.left == nil 和node.right == nil
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	fmt.Println(root)

	//nodes := []treeNode {
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)
	root.print()
	fmt.Println()
	root.left.right.setValue(4)
	root.left.right.print()
	pRoot := &root
	pRoot.setValue(300)
	pRoot.print()
	root.print()

	fmt.Println("\n遍历：")
	root.traverse()
}
// 指针接受者:改变内容；结构过大；一致性：有指针接受者就用指针接受者
// 值接收者是go特有
