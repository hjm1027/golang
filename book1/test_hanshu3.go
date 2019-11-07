//测试不用变量赋值能不能调用函数
package main

import "fmt"

func add(x int,y int) func() func() int {
		return func() func() int  {
				return func() int {
						return x+y
				}
		}
}


func main() {
		x,y:=15,16
		fmt.Println(add(x,y)()())
		A:=add(x int,y int)
		fmt.Println(A)//A和B输出了不同的地址，说明A和B是不同的指针，A指向了第一层函数，B指向了第二层（最后一层）函数，其中（）形式的参数列表是用来访问下一层函数的。
		B:=A()
		fmt.Println(B)
		C:=B()
		fmt.Println(C)
}
