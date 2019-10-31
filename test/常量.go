//数值常量
package main

import "fmt"

const (//数值常量也有类型，如果没有声明就按照上下文确定
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }//此处常量是int
func needFloat(x float64) float64 {
	return x * 0.1//此处常量是浮点64
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

