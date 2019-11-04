//测试make和指针传递有什么异同
package main

import "fmt"

func change(a,b int) {
		a=8
		b=4
}

func main() {
		var a,b = 5,3//传入指针才能保留处理，否则是单向传递
		change(a,b)
		fmt.Println(a,b)
}
