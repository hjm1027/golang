//range对索引和值的处理
package main

import "fmt"

func main() {
		pow := make([]int,10)
		for i := range pow {//不要值就把",value"删掉
				pow[i] = 1<< uint(i)
		}
		for _,value :=range pow {//不要索引赋值给_
				fmt.Printf("%d\n",valuei)
		}
}
//range将返回两个值，索引和值，带range的for循环会直接在切片或者数组遍历一遍
