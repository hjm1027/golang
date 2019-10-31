//测试哈希表在函数中的使用
package main

import "fmt"

func twoSum(nums []int, target int) /*[]int*/{

	var tmp =make([]int,2)

    var numMap map[int]int
    numMap = make(map[int]int)
	fmt.Println(numMap)//这个时候的哈希表还是什么东西都没有
    for i := 0; i < len(nums); i++ {
        var difference = target - nums[i]
        if v, ok := numMap[difference]; ok {
            //return []int={v,i}
			tmp[0]=v
			tmp[1]=i
        } else {
            numMap[nums[i]]=i
        }
    }
	fmt.Println(numMap)//这个时候的哈希表就有键值对了 推测是因为双赋值动了手脚
    //return []int{}
}

func main() {
		var nums =[]int{2,7,11,15}
		target := 9
		twoSum(nums,target)
		//fmt.Println(twoSum(nums,target)i)
}
