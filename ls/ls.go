package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		args  []string = os.Args
		dir   *os.File
		err   error
		Reset string = "\x1b[0m"
		Green string = "\x1b[32m"
		Blue  string = "\x1b[34m"
	)

	if len(args) > 1 {
		dir, err = os.Open(args[1])
	} else {
		dir, err = os.Open(".")
	}

	if err != nil {
		fmt.Println("Can't open directory:", err)
		return
	}
	defer dir.Close()

	info, err := dir.ReadDir(0)
	if err != nil {
		fmt.Println("Can't read directory:", err)
		return
	}

	for _, file := range info {
		if file.IsDir() {
			fmt.Print(Green , file.Name() , Reset, " ")
		} else {
			fmt.Print(Blue , file.Name() , Reset, " ")
		}
	}

	fmt.Println()
}
