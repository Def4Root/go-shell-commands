package main

import(
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get the current directory:", err)
		return
	}

	fmt.Println(dir)
}
