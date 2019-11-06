//输出参数索引和值，每行一个
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, value := range os.Args[1:] {
		fmt.Printf("%d %s\n", i, value)
	}
}
