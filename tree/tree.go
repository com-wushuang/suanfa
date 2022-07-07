package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *Node) {
	if root != nil {
		return
	}
	inOrder(root.Left)
	fmt.Print(root.Val, " ")
	inOrder(root.Right)
}

func postOrder(root *Node) {
	if root != nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Print(root.Val, " ")
}


