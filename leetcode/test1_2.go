//测试函数里面的哈希表是否自动存储了函数接收的参数
package main

import "fmt"

func test(num []int,target int)  {
		var m =make(map[int]int)//事实上函数里面建一个哈希表并不会自动存储接收的数值
		fmt.Println(m)
}


func main() {
		var num = make([]int,5)
		fmt.Println("请为数组赋值")
		for i:=0;i<5;i++{
				fmt.Scanf("%d",&num[i])
		}
		var target int
		fmt.Println("请输入目标数值")
		fmt.Scanf("%d",targeti)
		test(num,target)
}
