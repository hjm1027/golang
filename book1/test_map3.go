//测试map排序
//int类型
package main

import "fmt"

func main() {
		people := map[int]string{
				30:"Alice",
				18:"Jacksie",
				40:"Lucy",
				22:"Icy",
		}
		var ages []int
		for age := range people {
				ages = append(ages,age)
		}
		for i:=0;i<len(ages);i++{//冒泡排序法，第一个for遍历一遍，第二个for从i+1开始遍历
				for k:=i+1;k<len(ages);k++{
						if ages[i]>ages[k] {
								ages[i],ages[k]=ages[k],ages[i]
						}
				}
		}
		fmt.Println(ages)
		for _,age := range ages {
				fmt.Printf("%d %s  ",age,people[age])
		}
}
