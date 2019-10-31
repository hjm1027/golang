//结构体测试
package main

import "fmt"

type Vertex struct {
		X int
		Y int
}

func main() {
		v := Vertex{1,2}//这是直接结构体初始化，结构体变量类型的名字叫做Vertex
		v.X=4
		fmt.Println(v.X)//用'.'来访问结构体中的成员
		//fmt.Println(Vertex{1,2})
}

//结构体是定义了一种新的类型，要使用结构体就要先声明变量（结构体变量）
