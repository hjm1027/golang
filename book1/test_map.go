//测试不是map的键的特殊情况
package main

import "fmt"

func main() {
	name := map[string]int{
		"Alice":  30, //赋值后面要有逗号
		"Petter": 31,
	}
	name["string"]++
	fmt.Println(name["string"])
	fmt.Println(len(name))
}
