//测试map如果不申请内存，不初始化，能不能存值
package main

import "fmt"

func main() {
		var m =map[string]string{"China":"Beijing","America":"Washington"}//m后面的是声明类型，初始化的时候还要写一遍类型，这个时候前面的可以省略
		/*var m map[string]string
		m=make(map[string]string)
		m["China"]="Beijing"
		m["America"]="Washington"*/
		fmt.Println(m)
}
//只声明不初始化零值为nil，不能赋值。解决方法是用make创建申请内存。或者声明同时初始化两种。
