//切片
package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	fmt.Printf("%p%d\n",p,p)//切片名是一个指针,但是用‘*’也输出不了p所指的东西
}
