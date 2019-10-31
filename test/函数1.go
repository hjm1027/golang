//go中函数的实现
package main

import "fmt"

func add(x int,y int) int {//注意变量名int在变量后面
		return x+y
}

func main(){
		fmt.Println(add(42,13))//调用add函数，实参用逗号隔开
}
