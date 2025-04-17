package main

import (
	"fmt"

	"container/list"
)

func main() {

	//Linkedlist
	linkedList := list.New()

	//add to tail
	linkedList.PushBack(2)
	linkedList.PushBack(3)

	//add to head
	linkedList.PushFront(1)

	//iterate to Next
	print(*linkedList)

	//remove 2nd element
	fmt.Println("Removing 2nd element")
	second := linkedList.Front().Next()
	linkedList.Remove(second)

	print(*linkedList)

	//adding 2nd element
	fmt.Println("Adding 2nd element back")
	first := linkedList.Front()
	linkedList.InsertAfter(2, first)

	print(*linkedList)

}

func print(linkedList list.List) {
	curr := linkedList.Front()
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next()
	}
}
