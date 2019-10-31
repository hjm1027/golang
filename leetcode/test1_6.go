//自己用哈希表打一遍两数之和
package main

import "fmt"

func twoSum(nums []int,target int) []int {
		var m =make(map[int]int)
		var difference int
		for i:=0;i<len(nums);i++{
				difference = target-nums[i]
				if v,ok := map[difference]; ok{
						return []int{i,v}
				} else {
						map[nums[i]]=i//map的存储是反的，键是数组的值，存进去的是数组的索引。实在是太巧妙了。而v返回的是数组的索引，所以直接返回v就可以很方便了。
				}
		}
		return nil
}

func main() {
		var nums =[]int{2,7,11,15}
		target := 9
		fmt.Println(twoSum(nums,target))
}
