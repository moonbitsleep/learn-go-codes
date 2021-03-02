package main

import (
	"fmt"
	"learning-go/basics/newtree"
)

// 扩充别人类型的方法
type myNode struct {
	node *newtree.TreeNode
}

func (mine myNode) postOrder() {
	if mine.node == nil {
		return
	}

	myNode{mine.node.Left}.postOrder()
	myNode{mine.node.Right}.postOrder()
	mine.node.Print()
}

func main() {
	var root newtree.TreeNode
	root = newtree.TreeNode{Value: 3}
	root.Left = newtree.CreateNode(0)
	root.Right = newtree.CreateNode(5)
	root.Left.Right = &newtree.TreeNode{Value: 2}
	root.Right.Left = newtree.CreateNode(4)

	root.Traverse()
	fmt.Println()
	myNode{&root}.postOrder()
	fmt.Println()
}
