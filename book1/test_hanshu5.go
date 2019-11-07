//max函数求最大，用变长函数
package main

import "fmt"

func max(n ...int) int {
		if n==nil {//判断切片是否为空可以用长度是否为0，或者是否为零值（nil）但是不能用n[0]和nil比较，n[0]是int，但是nil是slice类型的零值。c里面nil也是指针的零值，索引类型的零值可能都是nil
				return 0
		}
		max:=n[0]
		for i:=1;i<len(n);i++{//判断还是要用三个表达式的循环,range循环不方便
				if n[i]>max {
						max = n[i]
				}
		}
		return max
}

func main() {
		fmt.Println(max())
}
