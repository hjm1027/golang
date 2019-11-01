//echo2的示例
package main

import (
		"fmt"
		"os"
)

func main() {
		count := 0
		s,sep:="",""
		for _,arg :=range os.Args[1:] {
				s +=sep + arg
				sep = ""
				count++
		}
		fmt.Println(s)
		fmt.Println(count)//由此可知循环是一遍也没有进行
}
