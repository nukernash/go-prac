package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Println(s)
}

func main() {

	go print("Hello World")

	fmt.Println("Hi")
	time.Sleep(time.Second)
}
