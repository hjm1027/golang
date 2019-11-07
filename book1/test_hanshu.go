//测试一个函数，是不是有一个多的布尔返回值
package main

import "fmt"

func add(x int,y int) int {
		return x+y
}

func main() {
		var x,y int
		fmt.Scanf("%d%d",&x,&y)
		x/*,y*/ = add(x,y)//说明还是不行的，多少个返回值就返回多少个值。
		fmt.Println(x)
}
