//测试string包newreader作用
package main

import (
    "fmt"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    var content ="hello,world.And hello you too"
    i:=strings.NewReader(content)
    fmt.Println(i)
    fmt.Printf("%T",i)
    dom,_:=goquery.NewDocumentFromReader(i)
    fmt.Println(dom)
}
