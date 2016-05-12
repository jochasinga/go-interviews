package main

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func main() {
	a := &LinkedListNode{Value: 5}
	b := &LinkedListNode{Value: 1}
	c := &LinkedListNode{Value: 9}

	a.Next = b
	b.Next = c
}
