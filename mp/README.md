# go-makeprogram

`mp [filename].[extension]` で拡張子に合うファイルを作成しテンプレートを記述するコマンドです。
mpはmakeprogramの略

## Usage
```
$ mp [filename].[extension]
```
これにより拡張子に合うファイルが作成され、テンプレートが記述されます。

## Example
```
$ mp main.go
```
main.go
``` go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
```

## Languages
- C
- Go
- HTML
- php