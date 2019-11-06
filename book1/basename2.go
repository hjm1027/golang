//简化版利用库函数strings.LastIndex
package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") //如果没找到“/”就会返回-1 找到了返回它所在的索引。
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(basename(s))
}
