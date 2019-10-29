//test2的另一种形式
package main

import "fmt"

func add(x,y int) int {//当函数形参多个相同时，可以写成这种形式
		return x+y
}

func main() {
		fmt.Println(add(42,13))
}
