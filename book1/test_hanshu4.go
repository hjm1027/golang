//测试变长函数的参数能不能为0
package main

import "fmt"

func add(n ...int) int {
		sum :=0
		for _,i := range n {
				sum+=i
		}
		return sum
}

func main() {
		fmt.Println(add())
}
