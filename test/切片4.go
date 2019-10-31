//测试切片的零值  slice的零值是'nil'
package main

import "fmt"

func main() {
		var z []int//切片的声明其实和数组的差不多，区别在于[]中为空
		fmt.Println(z,len(z),cap(z))
		if z == nil {//用'nil'去判断是否为空
				fmt.Println("nil")
		}
}
//切片和数组非常相似，可以说，切片是数组的一种形式。切片是引用，而数组是拷贝，改变切片，原来数组或原来切片中的元素也会发生改变。
