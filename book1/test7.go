//测试slice两次引用后改变元素的变化
package main

import "fmt"

func main() {
		array :=[5]int{1,2,3,4,5}
		s1 := array[0:3]
		s2 := s1[0:2]
		s2[0]=100
		fmt.Println(array,s1,s2)
}
