//测试数组的初始化
package main

import "fmt"

func main() {
		var a[10] int = [10]int{0,1,2,3,4,5,6,7,8,9}//int前面的中括号数字只能是数组元素个数，大括号里面的可以没有十个元素，少的用0填上
		var b[5] int = [5]int{0:3,1:6,4:3}//指定元素初始化，但是要全部都这种形式
		for i:=0;i<10;i++{
			fmt.Println(a[i])
	}
	for i:=0;i<5;i++{
			fmt.Println(b[i])
	}
	fmt.Println(a)//通过println输出a数组的指针a可以直接输出整个数组
	fmt.Printf("%p\n",a)//go中的数组名应该是指针，但是不知道为什么输出不了????
}//切片名肯定是指针
//初始化大括号赋值只能在初始化时进行，在运行过程中改变元素要使用for循环或者直接对相应元素赋值
