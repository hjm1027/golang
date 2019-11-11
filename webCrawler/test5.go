//测试网页搜索
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    //var keyword string
    //fmt.Scanf("%s",&keyword)
    rp,err := http.Get("https://search.bilibili.com/all?keyword=%E5%A4%A9%E6%B0%94%E4%B9%8B%E5%AD%90&from_source=banner_search")
    if err != nil {
        panic(err)
    }
    body,err:=ioutil.ReadAll(rp.Body)
    if err != nil {
        panic(err)
    }
    defer rp.Body.Close()

    dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        panic(err)
    }
    ele.Find("title").Each(func(i int,s *goquery.Selection){
        fmt.Println(s.Text())
    })
}

