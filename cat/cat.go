package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the filename.")
		return
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading the file.", err)
		return
	}

	fmt.Println(string(content))
}
