//阶乘和
package main

import (
		"fmt"
)

func jiecheng(n int) int {
		if n==0 {
				return 1
		} else {
				return n*jiecheng(n-1)
		}
}

func main() {
		var n int
		var sum int
		fmt.Scanf("%d",&n)
		for i:=1;i<=n;i++{
				sum+=jiecheng(i)
		}
		fmt.Println(sum)
}
