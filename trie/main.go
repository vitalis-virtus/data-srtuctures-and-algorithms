package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(w string) {
	var wordLength = len(w)
	var currentNode = t.root
	for i := 0; i < wordLength; i++ {
		var charIndex = w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	var wordLength = len(w)
	var currentNode = t.root

	for i := 0; i < wordLength; i++ {
		var charIndex = w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie)
	testTrie.Insert("toy")
	fmt.Println(testTrie.Search("toy"))
	fmt.Println(testTrie.Search("weapon"))
	fmt.Println(testTrie.Search("toys"))
}
