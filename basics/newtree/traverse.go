package newtree

// 分散一个方法在一个文件中
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	// 因为go语言允许nil传进去，所以不必判断node.left == nil 和node.right == nil
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
