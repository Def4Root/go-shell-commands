package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	file, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}