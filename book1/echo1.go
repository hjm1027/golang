//echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //声明了s和sep两个字符串变量 声明同时可以初始化，如果不初始化就会隐式初始化为这个类型的零值 数字的零值是0，string的零值是”“（空字符串）
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] //string类型的+运算是追加，把两个字符串连接在一起
		sep = ""
	}
	fmt.Println(s)
}

//:=是短变量声明，只能用在函数内部。并且可以同时声明多个变量。
//i++/i--是自增自减语句。这些都是语句，而不像c那样是表达式。所以j=i++是错的。应该要用j=i+1而且只有后缀，++i，--i都是错的。
