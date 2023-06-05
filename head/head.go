package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")
	for i := 0; i < 10 && i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}