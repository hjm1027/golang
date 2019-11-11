//随便从百度扒东西
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "strings"
)

func main() {
    var keywords string
    fmt.Scanf("%s",&keywords)
    rp,err:=http.Get("https://www.baidu.com/s?wd="+keywords+"&rsv_spt=1&rsv_iqid=0x92a1c4e500175cf7&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug3=4&rsv_sug1=5&rsv_sug7=100&rsv_sug2=0&inputT=2571&rsv_sug4=2571")
    fmt.Printf("%s\n",rp)
    if err!=nil {
        fmt.Println("false")
        panic(err)
    }
    body,err:=ioutil.ReadAll(rp.Body)
    defer rp.Body.Close()//引用要大写
    fmt.Printf("%s\n",body)
    dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        fmt.Println("false")
        panic(err)
    }
    dom.Find("a").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
