//测试make后的slice能否保存函数操作
package main

import "fmt"

func change(s []int) {
		s[0]=100
}

func main() {
		s :=make([]int,3)
		s = []int{1,2,3}
		change(s)
		fmt.Println(s)
}
