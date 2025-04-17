package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Hello World")
	testFile := "data/sample.txt"

	// Read a file content and print
	data, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Println("error while reading the file")
		return
	}

	fmt.Println(string(data))


	// Open a file and read line-by-line print
	file, err := os.Open(testFile)
	if err != nil {
		fmt.Println("error while opening the file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for(scanner.Scan()) {
		fmt.println(Strings.  "Line %d : %s", i , scanner.Text())
		i++;
	}
}
