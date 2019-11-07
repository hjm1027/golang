//用递归实现阶乘
package main

import "fmt"

func jiechen(n int) int {
		if n==0 {
				return 1
		} else {
				return n*jiechen(n-1)
		}
}

func main() {
		var n int
		fmt.Scanf("%d",&n)
		fmt.Println(jiechen(n))
}

