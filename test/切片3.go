//构造slice
//用make构造切片，make([]T,size,cap)一共有三个参数，其中第三个参数可以省略
//第二个表示长度，如果省略cap，那么申请的空间即为长度。cap表示申请的空间
//为什么要申请内存？因为要用到动态切片就要申请内存才能用，append
package main

import "fmt"

func main() {
		a := make([]int,5)
		printSlice("a",a)
		b := make([]int,0,5)
		printSlice("b",b)
		c := b[:2]
		printSlice("c",c)
		d := c[2:5]
		printSlice("d",d)
}

func printSlice(s string,x []int) {//函数定义在后面
		fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)//len()函数求切片的长度，cap（）用来求切片的空间。 %v返回相应值的默认格式感觉就是什么都可以，前提是默认值是你想要的。
}

