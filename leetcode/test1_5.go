//用数组实现哈希表解决两数之和的思路
//用数组实现哈希表的思想，但是会非常复杂，而且当要输出的数有多个的时候，因为每次都要遍历一遍存进去的数组，for循环的存在，存入数据会变得更加复杂。所以哈希表最厉害的地方，在于不用for循环遍历而可以直接通过range找出来又没有这个元素。
//顺便range也是可以遍历数组的，但是不能像map那样直接指定元素查找。返回索引和值两个数据。而map是返回指定元素的值和是否存在的false，所以我们可以直接用if去判断是否存在
package main

import "fmt"

func twoSum(nums []int,target int) []int {
		var s =make([]int,len(nums))
		var difference int
		for i:=0;i<len(nums);i++{
				difference = target - nums[i]
				for k:=0;k<len(s);k++{
						if s[k]==difference{
								return []int{i,k}
						}
				}
				s[i]=nums[i]
		}
		return nil
}


func main() {
		var nums =[]int{2,7,11,15}
		target :=9
		fmt.Println(twoSum(nums,target))
}
