//实现删除切片元素(其实只是把要删除的元素左右两边的切片拼起来了，并不存在真正的删除）
package main

import "fmt"

func main() {
		var i =[]int{1,2,3,4,5}
		var s =[]int{6,7,8,9}
		var a =[]int{10,11,12,13}
		//实现切片拼接
		s=append(i,s...)//在第二个后面打点实现拼接
		s=append(s,a...)
		fmt.Println(s)
		//实现切片元素删除
		index := 2
		i=append(i[:index],i[index+1:]...)
		fmt.Println(i)
}
