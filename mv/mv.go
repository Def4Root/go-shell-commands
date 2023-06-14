package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("please enter several arguments.")
		fmt.Println("Usage:", args[0], "[File] [destFile]")
		return
	}

	err := os.Rename(args[1], args[2])
	failOnError("failed to rename file", err)
}

func failOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s%s\n", msg, err)
	}
}
