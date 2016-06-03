// TODO: Write functions to automatically load a set of words into the tree
//
// Example from http://stackoverflow.com/questions/14708134/what-is-the-difference-between-trie-and-radix-trie-data-structures
//
//               *
//	        /
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
	"errors"
	"fmt"
	"strings"
)

/*
type Tree struct {
	
}

func NewTree(words ...string) (*Tree, error) {
	firstLetter := words[0][0]
	for _, word := range words {
		if word[0] != firstLetter {
			return nil, errors.New("Words do not share the same first letter.")
		}
	}
}
*/

type Edge struct {
	From  *Node
	To    *Node
	Label string
}

type Node struct {
	Name  string
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

				for {
					fmt.Printf("node %q is leaf: %t\n", next.Name, next.IsLeaf())
					if !next.IsLeaf() {
						_ = CheckWordInTree(suffix, next)
					} else {
						break
					}
					break
				}
				if strings.Compare(characters, init) == 0 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// Manually, laborously creating each node
	// nodes
	rootNode := &Node{Name: "root"}
	hNode := &Node{Name: "h"}
	elloNode := &Node{Name: "ello"}
	aNode := &Node{Level:2, Name: "a"}
	veNode := &Node{Level:3, Name: "ve"}
	tNode := &Node{Level:3, Name: "t"}
	
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
//	fmt.Println(CheckWordInTree("had", rootNode))
//	fmt.Println(CheckWordInTree("have", rootNode))
//	fmt.Println(CheckWordInTree("haven", rootNode))

	// Edge cases
	fmt.Println(CheckWordInTree("hatredenemy", rootNode))
}
