//关于结构体自己的测试
package main

import "fmt"

type student struct {
		name string
		num int
		mark int
}

func main() {
		var(
				stu1=student{"Jacksie",2019214228,100}
				stu2=student{name:"Kiki",num:2333}//如果指明了成员，那么就都要指明，不可以混用
		)
		stu2.mark = 99//一个var声明里面一个变量只能定义一次，否则重复定义
		fmt.Println(stu1,stu2)
}
//和c不同，go可以直接输出结构体，而c只能一个一个输出其中的成员
//还有go的结构体成员定义非常方便，声明同时可以进行定义，不定义就都是初始值
