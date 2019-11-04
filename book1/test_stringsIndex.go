//测试LastIndex函数返回值
package main

import (
		"fmt"
		"strings"
)

func main() {
		var s = "abcdefggfedcba"
		f := strings.LastIndex(s,"cba")//说明函数是找字符串中最后一个指定字符串并且返回该指定字符串第一个字符的索引
		fmt.Println(f)
}
