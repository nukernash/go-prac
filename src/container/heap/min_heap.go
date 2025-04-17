package main

import (
	"container/heap"
	"fmt"
)

type Person struct {
	name string
	age  int
}

type MinHeap []Person

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].age < h[j].age
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Person))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	len := h.Len()
	item := old[len-1]
	*h = old[:len-1]
	return item
}

func main() {
	avi := Person{
		name: "Avi",
		age:  35,
	}
	neel := Person{
		name: "Neel",
		age:  34,
	}
	h := &MinHeap{}

	heap.Init(h)

	heap.Push(h, avi)
	heap.Push(h, neel)

	p := heap.Pop(h)

	fmt.Println(p.(Person).name)

}
