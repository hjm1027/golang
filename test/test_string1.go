//测试string切出来的是不是切片
package main

import "fmt"

func main() {
    s:="hello,world"
    n:=s[1:4]//说明这个切出来的不是切片而是一个字符串
    i:=[]byte(n)
    i = append(i,'j')
    fmt.Printf("%s",i)
}
