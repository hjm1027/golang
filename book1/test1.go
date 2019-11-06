//测试短变量声明表达式要求是否类型一致
//测试说明可以不一致
package main

import "fmt"

func main() {
	a, b := 1, true
	fmt.Println(a, b)
}
