//爬取go中文网
//一页中的所有文档的标题。格式为: 第n个文档标题：title
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    num :=1
    rp,err:=http.Get("https://studygolang.com/topics")
    if err != nil {
        panic(err)
    }
    fmt.Println("hello")
    body,err := ioutil.ReadAll(rp.Body)
    if err != nil {
        panic(err)
    }
    defer rp.Body.Close()
    dom,err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        panic(err)
    }
    fmt.Println("hello")
    dom.Find("title").Each(func(i int,title*goquery.Selection) {
        fmt.Printf("%3d",num)
        fmt.Println(title.Text())
        num++
    })
}
