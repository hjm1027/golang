//测试切片的比较
package main

import "fmt"

func main() {
    a:=[]byte{'1','2','3','4'}
    b:=[]byte{'1','2','3','4'}
    if a == b{//这是错的，因为切片只能和nil比较，但是slice之间不能比较
        fmt.Println("ok")
    }
}
