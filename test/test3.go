//go语言函数可以同时返回任意数量个返回值
package main

import "fmt"
//构建swap函数去实现字符交换操作
func swap(x,y string) (string,string){//第二个括号表示返回值的数量和类型
		return y,x
}

func main() {
		a,b :=swap("hello","world")
		fmt.Println(a,b)
}
