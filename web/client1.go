//第一个爬虫，实现访问test5.go的服务端
package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    //这个里面一定要有http前缀，因为是通过http协议传输的，即使localhost也要。
    req,err := http.NewRequest("POST","http://localhost:9090/b",nil)
    if err != nil {
        panic(err)
    }
    res,err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    body,err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    dom,err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        panic(err)
    }
    dom.Find("body").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
