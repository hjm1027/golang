//用双指针删除重复元素
package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums)==0 {
        return 0
    }
    var i,k int
    i =0
    for k=1;k<len(nums);k++ {
            if nums[i]!=nums[k] {
                i++
                nums[i]=nums[k]
            }
        }
    return i+1
}

func main() {
		nums :=[]int{1,1,2,2,2,3,3,3,3,4,4}
		len := removeDuplicates(nums)
		fmt.Println(nums[:len])
}
