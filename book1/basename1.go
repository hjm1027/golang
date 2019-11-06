//移除路径和后面的.（只留一个）
package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	var s string
	fmt.Scanf("%s", &s) //说明string类型可以被重新赋值，也可以用scanf接收，注意是%s
	fmt.Println(basename(s))
}
