//实现动态数组 就是输入多少个，有多少个元素
package main

import "fmt"

func main() {
		var s =make([]int,1)
		var a int
		fmt.Println("请为数组赋值")
		for i:=0;i<10;i++{
				fmt.Scanf("%d",&a)//需要用一个另外的变量来存储输入的元素，然后再用append为切片扩容。
				s = append(s,a)//用append时还是要给原来的数赋值
		}
		fmt.Println(s)
}

