package newtree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

// node的方法
func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

// 传入指针，才能改变值
func (node *TreeNode) SetValue(v int) {
	node.Value = v
}

func CreateNode(value int) *TreeNode {
	// 返回局部变量的地址，不会出错
	//分配到堆上面，自动回收
	return &TreeNode{Value: value}
}


