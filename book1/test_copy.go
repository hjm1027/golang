//测试copy的特殊案例
package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := make([]int, 2)
	a = []int{1, 2, 3, 4, 5}
	b = []int{1, 2}
	c := copy(b, a)
	fmt.Println(b, c)
}
