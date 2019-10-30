//for循环中的range循环
package main

import "fmt"

var pow = []int{1,2,4,8,16,32,64,128}

func main() {
		for i,v :=range pow {//i表示索引，v是值
				fmt.Printf("2**%d = %d\n",i,v)
		}
}
