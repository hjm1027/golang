//移除元素(待修改）
package main

import "fmt"

func a(nums []int,val int) int {
    count:=len(nums)
    for i,n := range nums {
        //if i==0&&n==val
        if i==len(nums)-1&&n==val{
                count--
                break
            } else{
                if n==val{
                nums=append(nums[:i],nums[i+1:]...)
                count--
            }
            }
        }
        return count
}

func main() {
    var nums =[]int{3,3,2,1}
    val := 3
    length := a(nums,val)
    fmt.Println(length)
    nums = nums[:length]
    fmt.Println(nums)
}

