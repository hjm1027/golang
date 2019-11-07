//测试匿名函数
package main

import "fmt"

func main() {
		a := func(n1 int,n2 int) int {
				return n1+n2
		}
		fmt.Println(a(20,15))//没有参数列表会输出函数的地址。如果有参数列表但没有填写会报错。因为函数变量是引用类型。
}
