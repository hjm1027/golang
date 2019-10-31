//两数之和 输入一串数组，输入目标数值，遍历数组找两数之和是否相等 
//相等的组成数组输出
package main

import "fmt"

func twoSum(nums []int,target int) []int {
		var array =make([]int,2)
		for i:=0;i<len(nums);i++{
				for k:=i+1;k<len(nums);k++{
						if target==nums[i]+nums[k]{
								array[0]=i
								array[1]=k
								return array
						}
				}
		}
		return nil//函数最后一定要有return，即使没有返回任何值。如果在前面已经返回，那么这里就返回这个函数需要返回类型的零值
}

func main() {
		var nums []int=[]int{2,7,11,15}
		target :=9
		fmt.Println(twoSum(nums,target))
}
