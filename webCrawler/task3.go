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
    var cookie string ="ll=\"118254\"; bid=lEtJZhiJtxs; __utmc=30149280; __yadk_uid=OPRC2Mj3QDkp0O4R6AgmOvAbLOpKv4mr; push_noty_num=0; push_doumail_num=0; __utmv=30149280.20662; douban-profile-remind=1; __utmz=30149280.1573704403.5.3.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/passport/login; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1573707774%2C%22https%3A%2F%2Faccounts.douban.com%2Fpassport%2Flogin%22%5D; _pk_ses.100001.8cb4=*; __utma=30149280.1952350357.1573642118.1573704403.1573707775.6; ap_v=0,6.0; __utmt=1; dbcl2=\"206624187:Z+eC2q//ZGw\"; ck=Tzan; _pk_id.100001.8cb4=cc2d231d519253ea.1573642113.6.1573708649.1573704402.; __utmb=30149280.14.10.1573707775"

func getUrlRespHtml() string {
    url1 := "https://www.douban.com/people/206624187/"
    client:=&http.Client{}//创建一个client类型(客户端)变量
    req,err := http.NewRequest("GET",url1,nil)//通过NewRequest函数发送请求，第一个参数是发送方法，第二个是网址，第三个不知道是什么（暂时）返回响应报文要确定是否为客户端登陆，存储在req里面。
    if err != nil {
        fmt.Println("获取地址错误")
        panic(err)
    }

    req.Header.Set("Cookie",cookie)//在req报文头部加上cookie（用来确定客户端的一段信息）req应该是一种变量类型，Header是方法。
    req.Header.Add("User-Agent","Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")//在头部加上Agent，用来身份验证。
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
