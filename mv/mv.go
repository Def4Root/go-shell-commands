package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("please enter several arguments.")
		fmt.Println("Usage:", args[0], "[File] [destFile]")
		return
	}

	err := func() error {
		Path := args[1]
		destPath := args[2]

		File, err := os.Open(Path)
		failOnError("failed to open file:", err)
		defer File.Close()

		destFile, err := os.Create(destPath)
		failOnError("failed to create file:", err)
		defer destFile.Close()

		_, err = io.Copy(destFile, File)
		failOnError("failed to copy file:", err)

		err = os.Remove(Path)
		failOnError("failed to redmove file:", err)

		return nil
	}()

	failOnError("", err)
}

func failOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s%s\n", msg, err)
	}
}
