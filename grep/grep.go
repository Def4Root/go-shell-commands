package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	YELLOW = "\x1b[33m"
	RESET  = "\x1b[00m"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ", os.Args[0], "[file]", "[text]")
		return
	}

	source, err := ioutil.ReadFile(os.Args[1])
	failOnError("failed to read file", err)

	result := highlight(string(source), os.Args[2])
	fmt.Println(result)
}

func highlight(source string, text string) string {
	lines := strings.Split(source, "\n")
	var matches []string

	for _, line := range lines {
		if strings.Contains(line, text) {
			line = strings.ReplaceAll(line, text, YELLOW+text+RESET)
			matches = append(matches, line)
		}
	}

	return strings.Join(matches, "\n")
}

func failOnError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(0)
	}
}
