//实现切片的复制
//copy()函数  copy(destSlice,srcSlice []T) int
//srcSlice为数据来源切片 destSlice为复制目标，目标切片必须分配过空间且足够承载复制的元素个数。来源和目标的类型一致，copy的返回值表示实际发生复制的元素个数。srcSlice可以是一段，也可以是整个：copy(s,i[3:5])||copy(s,i)
package main

import "fmt"

func main() {
		var s1 =[]int{1,2,3,4,5}
		var s2 =[]int{6,7,8,9}
		copy(s1,s2[0:3])
		fmt.Println(s1)
}
