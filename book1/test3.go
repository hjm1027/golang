//测试make后在函数中是否双向传递
package main

import "fmt"

func change(a []int) {
	a[0] = 100
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	fmt.Println(a)
	a = make([]int, 5) //分配内存之后会归零值
	fmt.Println(a)
	change(a)
	fmt.Println(a) //确实分配内存之后也可以保留处理
}
