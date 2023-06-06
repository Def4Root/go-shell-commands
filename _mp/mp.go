package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("Usage:", arg[0], "<file>")
		return
	}

	extension := filepath.Ext(arg[1])

	var (
		gofmt string = "package main\n\n" + "import (\n" + "	\"fmt\"\n)\n\n" +
			"func main() {\n" + "	fmt.Println(\"Hello, World!\")\n" + "}"
		cfmt string = "#include <stdio.h>\n\nmain() {\n" +
			"    printf(\"Hello World\");\n" + "}"
		phpfmt  string = "<?php\n\n?>"
		htmlfmt string = "<!DOCTYPE html>\n<html lang=\"ja\">\n<head>\n" +
			"    <meta charset=\"UTF-8\">\n" +
			"    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n" +
			"    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n" +
			"    <title></title>\n</head>\n<body>\n\n</body>\n</html>\n"
	)

	file, err := os.Create(arg[1])
	if err != nil {
		fmt.Println("Can't create the file:", err)
		return
	}

	switch extension {
	case ".c":
		_, err = file.WriteString(cfmt)
	case ".go":
		_, err = file.WriteString(gofmt)
	case ".php":
		_, err = file.WriteString(phpfmt)
	case ".html":
		_, err = file.WriteString(htmlfmt)
	default:
		fmt.Println("Unsupported file extension:", extension)
		return
	}

	if err != nil {
		fmt.Println("Can't write to the file:", err)
		return
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Can't close the file:", err)
		return
	}
}
