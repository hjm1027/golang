//测试切片是引用的
package main

import "fmt"

func main() {
		var s1 =[]int{0,1,2,3,4,5,6,7,8,9}//s1后面可以没有int[]，但是赋值后面一定有[]int
		var s2 =make([]int,5)
		s2 =s1[1:6]
		fmt.Println(s2)
		s2[1]=3
		fmt.Println(s2)
		fmt.Println(s1)//证明是引用，改变切片的元素原来也会变
		s2=append(s2,2,3,3,7,8,9,10,11)
		fmt.Println(s2)
		fmt.Println(s1)//当切出来的切片扩充后长度在原来的范围内（是从对应的元素开始）切片添加的元素会覆盖掉原来切片的元素。但是当超过了这个范围，此操作无效，不会对原来的切片造成影响
}
