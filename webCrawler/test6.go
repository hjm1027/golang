//爬豆瓣网首页别人发的文章
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

var cookie string ="ll=\"118254\"; bid=0LNvjfUx_W4; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1573563221%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DYpwQMffJEsWWgWIck81T-ohlIOix7Ybl5lJdjOzJavK%26wd%3D%26eqid%3Dc987689000098893000000035dca4fa1%22%5D; _pk_id.100001.8cb4=c7a02bb4384b122e.1573370762.7.1573563549.1573559451.; __utma=30149280.366677727.1573372537.1573559142.1573563224.7; __utmz=30149280.1573539752.4.3.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmc=30149280; dbcl2=\"206624187:ZfxJoGAtWH0\"; ck=kBTR; __yadk_uid=kenKxbEsrfnrnUU9mvUm0hkMM0vmUrQM; push_noty_num=0; push_doumail_num=0; __utmv=30149280.20662; douban-profile-remind=1; ap_v=0,6.0; __gads=Test; _pk_ses.100001.8cb4=*; __utmb=30149280.4.10.1573563224; __utmt=1"

func getUrlRespHtml() string {
    url1 := "https://www.douban.com/note/739111474/"
    client := &http.Client{}
    req,err :=http.NewRequest("Post",url1,nil)
    if err != nil {
        fmt.Println("获取地址失败")
        panic(err)
    }
    req.Header.Set("cookie",cookie)
    req.Header.Set("User-Agent","Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")
    resp,err := client.Do(req)
    if err != nil {
        fmt.Println("登陆错误")
        panic(err)
    }
    resp_byte,err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    respHtml := string(resp_byte)
    return respHtml
}

func main() {
    fmt.Println(getUrlRespHtml())
    dom,err := goquery.NewDocumentFromReader(strings.NewReader(getUrlRespHtml()))
    if err != nil {
        panic(err)
    }
    dom.Find("legend").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
