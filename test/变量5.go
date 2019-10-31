//测试能不能同时赋值
package main

import "fmt"

func main() {
		//var a,b int= 10,是不能同时给a，b赋值的，会报错
		a,b:=10,20//这种形式却是可以的
		fmt.Println(a,b)
}
