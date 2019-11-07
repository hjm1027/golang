//go的删除重复项独特做法
package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i:=len(nums)-1; i>0; i-- {//通过判断i是否大于0来执行，这样i-1最小也是0
		if nums[i] == nums[i-1] {//不从前往后是因为你不知道长度，最后很难限定
			nums = append(nums[:i], nums[i+1:]...)//用...实现两个切片拼接。
		}
	}
	return len(nums)
	}

func main() {
		var nums =[]int{1,1,2,2,2,3,3,3,3,4,5}
		fmt.Println(nums[:removeDuplicates(nums)])
}
