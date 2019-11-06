//测试不同长度的数组单个元素能否比较
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [3]int{1, 2, 3}
	if a[0] == b[0] {
		fmt.Println("ok")
	}
}
