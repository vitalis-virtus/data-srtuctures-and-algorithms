package main

import "fmt"

type Stack struct {
	items []int
}

// push add a value to back
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// pop remove a value at the end
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	removed := s.items[l]
	s.items = s.items[:l]
	return removed
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}
