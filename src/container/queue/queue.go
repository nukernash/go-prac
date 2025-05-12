package main

import (
	"fmt"

	"container/list"
)

func main() {
	implUsingList()
	implUsingSlice()
}

func implUsingSlice() {
	queue := make([]int, 0)

	// enqueue
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)

	//Dequque
	front := queue[0]
	fmt.Println(front)
	queue = queue[1:]

	front = queue[0]
	fmt.Println(front)
	queue = queue[1:]

	//Peek
	front = queue[0]
	fmt.Println(front)

}

func implUsingList() {
	// queue
	stack := list.New()

	// Enqueue
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	// dequeue
	front := stack.Front()
	fmt.Println(front.Value)
	stack.Remove(front)

	// peek
	front = stack.Front()
	fmt.Println(front.Value)
}
