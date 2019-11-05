//测试make数组和指针传递是否能保存函数修改
package main

import "fmt"

func change(a *[5]int) {
		a[0]=100
}

func main() {
		array := [5]int{1,2,3,4,5}
		fmt.Println(array)
		change(&array)//说明数组名就是一个名字而不是一个指针,而且有一个问题，函数的参数很严格，规定了传入数组就是数组，而在c里面，传入数组一般都是通过传入指针来实现的。但这里并不是。所以传地址时，一定是&array  &array【0】是错的
		fmt.Println(array)
}
