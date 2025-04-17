package main

import (
	"fmt"
	"strings"
)

func process(s string, c chan string) {
	uppsercase := strings.ToUpper(s)

	c <- uppsercase
}

func main() {

	arr := []string{"hi there", "how are you", "hope you are doing well"}
	ch := make(chan string)

	for _, s := range arr {
		go process(s, ch)
	}

	for i := 0; i < len(arr); i++ {
		result := <-ch
		fmt.Printf("%d : %s\n", i, result)
	}

}
