//return没有返回值的函数（但是有返回的参数）
package main

import "fmt"

func split(sum int) (x,y int) {
		x=sum*4/9
		y=sum-x
		return
}

func main() {
		fmt.Println(split(17))
}

