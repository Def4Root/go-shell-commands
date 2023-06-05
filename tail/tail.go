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
		fmt.Println("Can't read file:", err)
	}

	lines := strings.Split(string(file), "\n")

	if len(lines) > 10 {
		lines = lines[len(lines)-10:]
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
