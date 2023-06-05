package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "dirname")
		return
	}

	err := os.Mkdir(os.Args[1], 0755)
	if err != nil {
		fmt.Println("Directory create error:", err)
		return
	}
}