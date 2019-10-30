//结构体测试2
package main

import "fmt"

type Vertex struct {
		X int
		Y int
}

func main() {
		v :=Vertex{1,2}
		p :=&v
		p.X = 1e9//表示1*10的九次方
		fmt.Println(v)//通过指针来访问
}
//通过指针间接的访问是透明的。？
