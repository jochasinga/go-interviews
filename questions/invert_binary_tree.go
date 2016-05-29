package main

import "fmt"

type TreeNode struct {
	Val string
	Left *TreeNode
	Right *TreeNode
}

func swapTree(n *TreeNode) *TreeNode {
	if n != nil {
		tmp := n.Left
		n.Left = swapTree(n.Right)
		n.Right = swapTree(tmp)
	}

	return n
}

func main() {
	// Haphazardedly create a simple tree
	
	root := &TreeNode{Val:"Root"}
	left1 := &TreeNode{Val:"L"}
	right1 := &TreeNode{Val:"R"}
	
	left2 := &TreeNode{Val:"LL"}
	right2 := &TreeNode{Val:"LR"}

	left3 := &TreeNode{Val:"RL"}
	right3 := &TreeNode{Val:"RR"}
	

	root.Left = left1
	root.Right = right1

	root.Left.Left = left2
	root.Left.Right = right2

	root.Right.Left = left3
	root.Right.Right = right3

	
	root = swapTree(root)
	
	fmt.Printf("First left: %q\n", root.Left.Val)
	fmt.Printf("Second left: %q\n", root.Left.Left.Val)
	fmt.Printf("Second right: %q\n", root.Left.Right.Val)

	fmt.Printf("First right: %q\n", root.Right.Val)
	fmt.Printf("Second left: %q\n", root.Right.Left.Val)
	fmt.Printf("Second right: %q\n", root.Right.Right.Val)
}
