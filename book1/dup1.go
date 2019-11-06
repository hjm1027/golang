//输出标准输入中出现次数大于1的行 前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) //这里声明规定了input在缓冲区读取
	for input.Scan() {                  //scan函数用来读取下一行，读不到时返回false，读到返回true。
		counts[input.Text()]++ //等价于line := input.Text() Text函数读取内容
	} //counts[line]=counts[line]+1 相当于存了行又计算了行出现的次数
	//键在map中不存在也是没有问题的，会根据类型推导直接初始化为零值。
	for line, n := range counts { //map中的range是随机的，每一次都不一样。
		if n > 1 { //这里是至少出现两次才会输出。
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
