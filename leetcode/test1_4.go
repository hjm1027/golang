//证实是把键代进去实现哈希表的自动存储
package main

import "fmt"

func twoSum(nums []int, target int) {
		var numMap = make(map[int]int)
		var difference = target-nums[2]
        v, ok := numMap[difference]
		fmt.Println(v,ok)
		fmt.Println(numMap)//事实证明，单纯放一个数进去找，是不能实现自动存储的
}

func main() {
		var nums []int=[]int{2,7,11,15}
		target := 9
		twoSum(nums,target)
}
//事实上哈希表不能实现自动存储，示例代码是检查哈希表里面有没有一个值相匹配，没有就把目前这个值存进去，然后再找下一个值判断是不是相匹配。如果是就输出，不是就继续存。构思非常巧妙。因为是一个数组，那么遍历过的都会存进去，相当于只遍历了一次就实现了。可以用数组也试一下。
