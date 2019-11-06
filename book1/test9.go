//测试长度不同数组能不能比较
package main

import "fmt"

func main() {
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [5]int{1, 2, 3, 4}
	if a1 == a2 {
		fmt.Println("ok")
	}
}
