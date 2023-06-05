package main

import (
	"fmt"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(fromKey, toKey int) {
	// get vertex
	fromVertex := g.getVertex(fromKey)
	toVertex := g.getVertex(toKey)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", fromKey, toKey)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, toKey) {
		err := fmt.Errorf("existing edge (%v-->%v)", fromKey, toKey)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.Vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex %v: ", v.key)

		for _, val := range v.adjacent {
			fmt.Printf("%v", val.key)
		}
	}
}

func main() {
	testG := Graph{}

	for i := 0; i < 5; i++ {
		testG.AddVertex(i)
	}

	testG.AddEdge(1, 2)

	testG.Print()
}
