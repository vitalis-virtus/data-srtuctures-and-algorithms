package main

import "fmt"

// MaxHeap has a slicethat holdsthe array
type MaxHeap struct {
	array []int
}

// insert adds element to heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// maxHeapify will heapify from bottom to top
// maxHeapifyUp accepts int when heapify should start
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

// get left child index
func left(index int) int {
	return 2*index + 1
}

// get right child index
func right(index int) int {
	return 2*index + 2
}

// swap keys in array
func (h *MaxHeap) swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func main() {
	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 5, 7, 16, 19, 25}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m.array)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
