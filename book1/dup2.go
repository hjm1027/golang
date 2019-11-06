//dup2 打印输入中多次出现的行的个数和文本
//它从stdin或者指定文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] //读取哪个文件在命令行输入
	if len(files) == 0 { //若命令行没有文件
		countLines(os.Stdin, counts) //就在缓冲区读取
	} else { //如果命令行里有文件名
		for _, arg := range files { //遍历获取文件名的值存进arg里面。
			f, err := os.Open(arg) //open函数返回两个值 第一个是打开的文件，第二个是此函数内置error类型的值 如果err等于特殊的内置nil值，表明文件打开成功
			if err != nil {        //不等于nil说明打开失败
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err) //给提示错误信息 然后继续循环
				continue                                 //继续循环打开下一个文件
			}
			countLines(f, counts) //输入函数通过scanner扫描，text获取存进map。
			f.Close()             //关闭文件
		}
	}
	for line, n := range counts { //遍历map，输出出现次数和每行内容
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//函数countLines声明 调用和声明不分先后

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	//逐行读取，所以是分流式的
	for input.Scan() {
		counts[input.Text()]++ //countLines函数把行数和出现次数都存进map里面。
	}
}

//这个版本的dup使用“流式”模式读取输入，然后按需拆分为行。这样原理上可以处理海量输入。还有一个方法是一次读取整个输入到大块内存，一次性地分割所有行。然后处理这些行。
