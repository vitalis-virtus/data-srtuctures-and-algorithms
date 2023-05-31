package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(150)
	tree.Insert(12)
	tree.Insert(153330)
	tree.Insert(15120)
	tree.Insert(1)
	tree.Insert(1520)
	tree.Insert(185110)
	tree.Insert(1750)
	tree.Insert(1560)

	fmt.Println(tree)
	fmt.Println(tree.Search(11))
}
