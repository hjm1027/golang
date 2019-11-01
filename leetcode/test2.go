//测试整型溢出
package main

import "fmt"

func reverse(x int) int {
		var y int
		for x != 0{//这个是while语句的实现。如果是for，那么应该是for ;(判断语句);{
				y=y*10+x%10//for要有两个分号,而while没有
				if x>1<<31-1||x<(-1)<<31{//负数一定要打括号，不然会报错。
						return 0//<<表示转换成二进制后向左移动，就是扩大2的n次方。
				}//>>表示二进制向右移动，缩小2的n次方
				x/=10
		}
		return y
}

func main() {
		var x int
		fmt.Scanf("%d",&x)
		fmt.Println(reverse(x))
}
//二进制运算符还有& | 和^，分别是二进制两个数值按位对比，如果都为1则为1，有一个是0则另一个变为0、如果都为0则为0，又一个为1则都变为1、如果相同位上数值不同（0，1）就变成1
