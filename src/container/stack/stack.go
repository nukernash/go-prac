package main

import (
	"fmt"

	"container/list"
)

func main() {

	implUsingSlice()

	implUsingContainerList()
}

func implUsingSlice() {

	//STack
	stack := make([]int, 0)

	//Push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)

	//peek
	front := stack[len(stack)-1]
	fmt.Println(front)

	// pop
	stack = stack[:len(stack)-1]
	next := stack[len(stack)-1]
	fmt.Println(next)

}

func implUsingContainerList() {
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
