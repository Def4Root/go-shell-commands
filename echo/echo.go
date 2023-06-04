package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var result []string

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "$") {
			envVar := os.Getenv(strings.TrimPrefix(arg, "$"))
			result = append(result, envVar)
		} else {
			result = append(result, arg)
		}
	}

	output := strings.Join(result, " ")
	fmt.Print(output, "\n")
}
