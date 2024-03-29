//comma函数 向表示十进制非负整数的字符串中插入逗号
package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(comma(s))
}
