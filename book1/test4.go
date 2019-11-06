//测试数组重新赋值
package main

import "fmt"

func main() {
	/*var array =[3]int{1,2,3}
	fmt.Println(array)
	array = [3]int{3,2,1}
	fmt.Println(array)
	array[0] = 4
	fmt.Println(array)
	var a = [...]int{1,2,3,4}
	fmt.Println(a)*/
	//type Currency int
	const (
		USD int = iota //iota常量只能在const声明中使用，初始化为0，每次赋值叠加1
		EUR
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "*", RMB: "%"}
	fmt.Println(RMB, symbol)
}
