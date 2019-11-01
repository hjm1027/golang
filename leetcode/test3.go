//回文数的实现
package main

import "fmt"

func isPalindrome(x int) bool {
		if x<0 {
				return false
		}
		var num =make([]int,0)
		len := 0
		for x!=0{
				num = append(num,x%10)//把整数存进数组里面，动态数组的实现
				x/=10
				len++
		}
		k := len-1
		for i:=0;i<len/2;i++{
				if num[i]!=num[k]{//在一个循环里面判断，另外一边通过变量k来实现
						return false
				}
				k--
		}
		return true
}//记得最后加上大括号

func main() {
		var num int
		fmt.Println("请输入一个整数")
		fmt.Scanf("%d",&num)
		fmt.Println(isPalindrome(num))
}
