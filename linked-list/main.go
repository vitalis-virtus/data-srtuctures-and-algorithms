package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 45}
	node2 := &node{data: 90}
	node3 := &node{data: 150}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	fmt.Println(myList)
}
