//测试结构体不定义结果
package main

import "fmt"

type student struct {
		name int
		num string
		mark int
}

func main() {
		stu1 :=student{}
		var stu2 =student{}
		fmt.Println(stu1,stu2)
}
//不管用哪种方式声明结构体变量，后面的结构体类型名student都要加上大括号
