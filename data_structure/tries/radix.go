// TODO: Improve
package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Level int
	Edge []string
	Next []*Node
}

func CheckWordInTree(word string, first *Node) bool {
	prefix := ""
	suffix := ""
	next := &Node{}
	wordLen := len(word)
	for i, _ := range word {
		prefix = word[:i+1]
		for _, e := range first.Edge {
			if strings.Contains(e, prefix) {
				suffix = word[i+1:wordLen]
				next = first.Next[0]
			}
		}
		for _, e := range next.Edge {
			if strings.Contains(e, suffix) {
				return true
			}
		}
	}
	return false
}


func main() {
	// "hello" route
	root := &Node{Level: 0, Edge: []string{"h"}}
	h := &Node{Level: 1, Edge: []string{"ello"}}
	ello_ := &Node{Level: 2, Edge: []string{"$"}}
	root.Next = append(root.Next, h)
	h.Next = append(h.Next, ello_)

	// "have" route
	h.Edge = append(h.Edge, "a")
	a := &Node{Level: 3, Edge: []string{"ve"}}
	ve_ := &Node{Level: 4, Edge: []string{"$"}}
	h.Next = append(h.Next, a)
	a.Next = append(a.Next, ve_)

	// "hat" route
	t_ := &Node{Level: 4, Edge: []string{"$"}}
	a.Next = append(a.Next, t_)
	a.Edge = append(a.Edge, "t")

	fmt.Println(CheckWordInTree("hello", root))
}
