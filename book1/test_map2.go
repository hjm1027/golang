//测试map的排序 用数组的形式
//键是字符串的情况：
package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"Chuchu":  18,
		"Petter":  32,
		"Jacksie": 18,
		"Kiki":    20,
		"Alice":  30,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)//这个函数就是以英文字母大小写顺序来排序的，abcd....
	fmt.Println(names)
	for _, name := range names {
		fmt.Printf("%s %d  ", name, ages[name])
	}
	fmt.Printf("\n")
}
