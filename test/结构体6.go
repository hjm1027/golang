//go语言中结构体数组的实现
package main

import "fmt"

type student struct {
		name string
		num int
		mark int
}

func main() {
		var a[3] student//student是我定义的一种新的类型，而用这种类型声明数组，这个数组就是结构体数组，每个元素包含三个成员
		//结构体数组的初始化
		for i:=0;i<3;i++{//与c是一样的
			fmt.Scanf("%s%d%d",&a[i].name,&a[i].num,&a[i].mark)
	}
		for i:=0;i<3;i++{
				fmt.Println(a[i])//只打印两个0，因为字符串的初始值是空字符串，显示不出来
		}
}
