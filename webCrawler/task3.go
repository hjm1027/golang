//模拟登陆爬豆瓣网
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

//把cookie复制变成全局变量
    var cookie string ="ll=\"118254\"; bid=0LNvjfUx_W4; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1573539751%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DYpwQMffJEsWWgWIck81T-ohlIOix7Ybl5lJdjOzJavK%26wd%3D%26eqid%3Dc987689000098893000000035dca4fa1%22%5D; _pk_id.100001.8cb4=c7a02bb4384b122e.1573370762.4.1573539773.1573535471.; __utma=30149280.366677727.1573372537.1573535477.1573539752.4; __utmz=30149280.1573539752.4.3.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmc=30149280; dbcl2=\"206624187:ZfxJoGAtWH0\"; ck=kBTR; __yadk_uid=kenKxbEsrfnrnUU9mvUm0hkMM0vmUrQM; push_noty_num=0; push_doumail_num=0; __utmv=30149280.20662; douban-profile-remind=1; ap_v=0,6.0; _pk_ses.100001.8cb4=*; __utmb=30149280.6.10.1573539752; __utmt=1"

func getUrlRespHtml() string {
    url1 := "https://www.douban.com/people/206624187/"
    client:=&http.Client{}//创建一个client类型(客户端)变量
    req,err := http.NewRequest("GET",url1,nil)//通过NewRequest函数发送请求，第一个参数是发送方法，第二个是网址，第三个不知道是什么（暂时）返回响应报文要确定是否为客户端登陆，存储在req里面。
    if err != nil {
        fmt.Println("获取地址错误")
        panic(err)
    }

    req.Header.Set("Cookie",cookie)//在req报文头部加上cookie（用来确定客户端的一段信息）req应该是一种变量类型，Header是方法。
    req.Header.Add("User-Agent","Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")//在头部加上Agent，用来破解反爬虫
    resp,err := client.Do(req)//通过Do方法从客户端发送请求，用resp存储响应报文
    if err != nil {
        fmt.Println("登陆错误")
        panic(err)
    }

    //client.Do发送请求后，后面的步骤就一样了。通过readall读取，然后编辑
    resp_byte,err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    respHtml := string(resp_byte)
    return respHtml
}

func main() {
    //fmt.Println(getUrlRespHtml())
    dom,err := goquery.NewDocumentFromReader(strings.NewReader(getUrlRespHtml()))
    if err != nil {
        panic(err)
    }
    dom.Find("a").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
