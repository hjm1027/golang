//采用一次性输入所有内容，再分割行处理
//只读取文件，不读取标准输入
//dup3示例
package main

import (
		"fmt"
		"io/ioutil"
		"os"
		"strings"
)

func main() {
		counts := make(map[string]int)
		for _,filename := range os.Args[1:] {
				data,err := ioutil.ReadFile(filename)//用readfile函数直接获取文件全部内容。
				if err != nil {
						fmt.Fprintf(os.Stderr,"dup3=%v\n",err)
						continue
				}
		for _,line := range strings.Split(string(data),"\n") {//split函数将一个字符串分割成一个由子串组成的slice 这里的子串就是每行的字符串
				counts[line]++
		}
}
for line,n := range counts {
		if n>1 {
				fmt.Printf("%d\t%s\n",n,line)
		}
}
}
//dup2分流式的读取是通过for scan（）扫描来循环一行一行遍历。而dup3的是通过把原文件分给成若干行再用range遍历。一开始都是接收了内容，但是open接收的应该是这个文件的指针。而readfile是真的接收了内容。还有前者一行一行读取，后者全部都取之后分开再每个读取。
