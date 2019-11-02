//完全平方数 加上100是完全平方 再加上168还是完全平方
//大概有两种最基本的思路
//第一种是这个数用sqrt开方，开方返回的是浮点数，通过四舍五入（+0.5）判断n*n?=原来哪个数
//第二种是在一个区间枚举（遍历实现），找有没有n的平方等于这个数。
//方法一
package main

import (
		"fmt"
		"math"
)

func main() {
		var a,b int
		for i:=0;;i++{
				a=int(math.Sqrt(float64(i+100))+0.5)//go没有隐式类型转换，要手动转
				b=int(math.Sqrt(float64(i+268))+0.5)//sqrt的参数和返回值都是float64
				if a*a==i+100&&b*b==i+268 {
						fmt.Println(i)
						break
				}
		}
}
