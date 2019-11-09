//原地修改元素
package main

import "fmt"

func removeElement(nums []int,val int) int {
    n:=len(nums)
    for i:=0;i<n;i++{
        if i==len(nums)-1&&nums[i]==val {
            n--
            break
        }
        if nums[i]==val {
            copy(nums[i:],nums[i+1:])
            n--
        }
    }
    return n
}

func main() {
    nums:=[]int{3,2,2,3}
    fmt.Println(removeElement(nums,3))
}
