//通过双赋值检测某个值的存在
package main

import "fmt"

type student struct{
		num int
		mark int
}

func main() {
		var stu map[string]student//这种结构体的写法和结构体有点像，但它却不能指定输出成员，要输出，就会全部一起输出
		stu=make(map[string]student)
		stu["Tom"]=student{2019214226,88}//结构体赋值要在大括号前面加上类型
		fmt.Println(stu["Tom"])
		v,ok :=stu["Jacksie"]//对map双赋值，v赋值的是对应的值，当值存在时，返回true给ok，不存在时返回false给ok
		fmt.Println("The value:",v,"Present?",ok)
}
