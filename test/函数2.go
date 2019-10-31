//函数的更简洁形式
package main

import "fmt"

func add(x,y int) int {//类型一样的可以省略前面的类型，只写最后一个
		return x+y
}

func main() {
		var x,y int//声明也是一样的
		fmt.Scanf("%d%d",&x,&y)//go的scanf是回车后接收完成，用空格分开接收。如果什么都不输入，那么默认值为0
		fmt.Println(add(x,y))
}
