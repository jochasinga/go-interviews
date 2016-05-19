// TODO: Search should search until hits the node whose
// Next is empty to make sure the word is completely terminated.
//
// Example from http://stackoverflow.com/questions/14708134/what-is-the-difference-between-trie-and-radix-trie-data-structures
//
//              *
//	       /
//         (ello)
//           /
//  * - h - * -(a) - * - (t) - *
//	              \
//                    (ve)
//                       \
//                        *
//
//
package main

import (
	"fmt"
	"strings"
)

type Edge struct {
	From *Node
	To *Node
	Label string
}

type Node struct {
	Level int
	Heads []*Edge
}

func (n *Node) IsLeaf() bool {
	if len(n.Heads) > 0 {
		return false
	}
	return true
}

func (n *Node) IsRoot() bool {
	if n.Level > 0 {
		return false
	}
	return true
}

var (
	counter = 0
	characters = ""
)


func CheckWordInTree(word string, node *Node) bool {

	// FIXME: Relies on a global counter in order
	// to persist the initial word provided in the argument.
	init := ""
	
	if counter == 0 {
		init = word
	}
	
	fmt.Printf("collected: %q\n", characters)
	fmt.Printf("%d RUN\n", counter)
	
	counter++

	for i, _ := range word {
		prefix := word[:i+1]

		for _, edge := range node.Heads {
			if strings.Compare(edge.Label, prefix) == 0 {
				characters += prefix
				suffix := word[i+1:len(word)]
				next := edge.To

				// FIXME: The recursion never hits the leaf node, resulting
				// in a forever loop if a finite match isn't found.
				for counter > 5 {
					_ = CheckWordInTree(suffix, next)
					if strings.Compare(characters, init) == 0 {
						return true
					}
				}

			}
		}
	}
	return false
}

func main() {
	// TODO: Let's write a function to load any set of words
	// into the tree automatically!
	
	// nodes
	rootNode := &Node{}
	hNode := &Node{}
	elloNode := &Node{IsLeafNode: true}
	aNode := &Node{Level:2}
	veNode := &Node{Level:3, IsLeafNode: true}
	tNode := &Node{Level:3, IsLeafNode: true}
	
	// create Edges
	hEdge := &Edge{
		From: rootNode,
		To: hNode,
		Label: "h",
	}
	elloEdge := &Edge{
		From: hNode,
		To: elloNode,
		Label: "ello",
	}
	aEdge := &Edge{
		From: hNode,
		To: aNode,
		Label: "a",
	}
	veEdge := &Edge{
		From: aNode,
		To: veNode,
		Label: "ve",
	}
	tEdge := &Edge{
		From: aNode,
		To: tNode,
		Label: "t",
	}
	
	rootNode.Heads = []*Edge{hEdge}
	hNode.Heads = []*Edge{elloEdge, aEdge}
	aNode.Heads = []*Edge{tEdge, veEdge}


//	fmt.Println(CheckWordInTree("hello", rootNode))
//	fmt.Println(CheckWordInTree("ha", rootNode))
//	fmt.Println(CheckWordInTree("hat", rootNode))
	fmt.Println(CheckWordInTree("had", rootNode))
//	fmt.Println(CheckWordInTree("have", rootNode))

	// Edge cases
//	fmt.Println(CheckWordInTree("hate", rootNode))
}
