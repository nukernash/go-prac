package main

import (
	"fmt"

	"container/list"
)

func main() {

	//queue
	stack := list.New()

	//Enqueue
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	//dequeue
	front := stack.Front()
	fmt.Println(front.Value)
	stack.Remove(front)

	//peek
	front = stack.Front()
	fmt.Println(front.Value)

}
