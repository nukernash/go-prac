package main

import (
	"fmt"
	file_processor "go-prac/src/concurrency/threadpool/unique_file_content/file_processor"
)

func main() {
	fp := &file_processor.FileProcessor{
		Root: "./test_data",
	}

	result := fp.ProcessSequentially()
	fmt.Printf("Unique File content processed sequentially : %+v\n", result)

	result = fp.ProcessConcurrently()
	fmt.Printf("Unique File content processed concurently : %+v\n", result)
}
