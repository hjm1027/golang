//测试成员顺序不同是否相等
//根本不存在上述情况，因为结构体只能相同类型比较，就好像不同类型不能比较，int和string不能比较，其实是一样的。
package main

import "fmt"

type student struct {
		num int
		mark int
		name string
}

type student2 struct {
		mark int
		num int
		name string
}

func main() {
		var s1 student
		var s2 student
		if s1==s2 {
				fmt.Println("hello")
		}
}
