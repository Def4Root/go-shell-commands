package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	lines := strings.Split(string(content), "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}
}