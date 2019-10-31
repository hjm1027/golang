//向切片中添加元素，通过append实现
package main

import "fmt"

func main() {
		var a []int
		printSlice("a",a)
		//append为一个空间和长度都为0的切片添加元素
		a = append(a,0)//添加元素0
		printSlice("a",a)
		//切片会自己扩充 当空间或长度不够时。而且通常以两倍形式，2、4、6、8、16
		a = append(a,1)
		printSlice("a",a)
		//还可以一次性添加多个元素
		a = append(a,2,3,4,5,6,7,8,9)//通常是两倍两倍扩张，当两倍之后还不够时，会变成扩张后数组+1。
		printSlice("a",a)
}

func printSlice(s string,x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}
