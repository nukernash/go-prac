package main

import (
	"fmt"

	"container/list"
)

func main() {

	//Stack
	stack := list.New()
	stack.PushFront(1)
	stack.PushFront(2)
	stack.PushFront(3)

	//pop
	front := stack.Front()
	fmt.Println(front.Value)
	stack.Remove(front)

	//peek
	front = stack.Front()
	fmt.Println(front.Value)

}
