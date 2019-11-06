//测试strings包里面join的功能
package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"12", "23", "34"}
	str := strings.Join(slice, "") //如果后面的是空字符串，那么可以实现最直观地连接。
	fmt.Println(str)
}
