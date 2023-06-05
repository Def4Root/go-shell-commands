package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir(os.Args[1], 0755)
	if err != nil {
		fmt.Println("Directory create error:", err)
		return
	}
}