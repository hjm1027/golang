//用函数闭包实现斐波那契数列
package main

import "fmt"

func fibonacci() func() int{//闭包是指内嵌函数输出时返回的值
		a:=-1//这里的a，b每次用参数引用，每次改变都会保留（有点像全局变量）
		b:=1//用不同参数引用会与其他引用隔离开（这是和全局变量不同的地方）
		c:=0
		return func() int {
				c=b//要用一个变量来存储新一轮迭代之前的前一个数
				b=a+b//也可用两个变量来写 a,b = b,a+b 同时赋值时，此处的a是赋值前的
				a=c
				return b
		}
}

func main() {
		f:= fibonacci()
		for i:=0;i<10;i++{
				fmt.Println(f())
		}
}
//用参数引用函数的意思就是f:= fibonacci()，这个f就是参数，:=就是引用。
