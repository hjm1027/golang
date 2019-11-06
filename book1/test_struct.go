//测试结构体的成员能不能是结构体
package main

import "fmt"

type point struct {
		x int
		y int
}

type car struct {
		p point//结构体里面可以放别的结构体，但是不能放自身
		fuel int
		//c car
}

func main() {
		var c car
		fmt.Println(c.p.x)
}
