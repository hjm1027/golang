//echo3  用join函数代替数组拼接处理
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "")) //用Join函数来实现字符串切片的拼接。这个拼接是切片内部的拼接，把多个字符串拼成同一个。slice={"12","23","45"}  strings.Join(slice,"--") 则结果输出为 ""12--23--45"
}
