//结构体文法
package main

import "fmt"

type Vertex struct {
		X,Y int
}

var (//可以用var小括号形式同时定义多个结构体
		v1 = Vertex{1,2}
		v2 = Vertex{X:1}//没有定义Y，所以输出了Y的初始值0
		v3 = Vertex{}//都没有定义所以两个都是0
		p  = &Vertex{1,2}//这是指针类型
)//大括号里面定义成员，指定成员要用X：冒号赋值，如果同时赋值直接逗号，隔开

func main() {
		fmt.Println(v1,p,v2,v3)
		fmt.Printf("%p\n",p)
}
