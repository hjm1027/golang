//用密码用户名形式模拟登陆
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "github.com/howeyc/gopass"
    "github.com/tidwall/gjson"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    requestUrl := "https://accounts.douban.com/j/mobile/login/basic"

    var name string
    fmt.Print("豆瓣账号：")
    _,_=fmt.Scanln(&name)//获取用户名
    fmt.Print("输入密码:")
    password,_ :=gopass.GetPasswdMasked()//密码加密获取更安全

    data := url.Values{}//声明一个values的类型
    data.Set("name",name)//用类型的方法进行设置，这一段是设置在了body里面通过formdata格式（表单）来传输。
    data.Set("password",string(password))
    data.Set("remember","false")
    data.Set("ck","")
    data.Set("ticket","")

    payload := strings.NewReader(data.Encode())//把data数据放在playload里

    req,err := http.NewRequest("POST",requestUrl,payload)//通过请求发送一次用户名信息，存储返回值
    if err != nil {
        panic(err)
        return
    }
//在返回信息里加上header头部
    req.Header.Add("Accept","application/json")
    req.Header.Add("Content-Type","application/x-www-form-urlencoded")
    req.Header.Add("Origin","https://accounts.douban.com")
    req.Header.Add("Referer","https://accounts.douban.com/passport/login_popup?login_source=anony")
    req.Header.Add("Sec-Fetch-mode","cors")
    req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
    req.Header.Add("X-Requested-with","XMLHttpRequest")

    //fmt.Printf("%+v",req)
    //fmt.Printf("\n\n")
//再通过客户端发送一次请求报文，响应报文存在res里面
    res,err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
        return
    }
    defer res.Body.Close()
//把res的body部分取出来
    body,err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
        return
    }
    fmt.Println(res.Status)

    var UserAgent ="Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36"

    result := gjson.Get(string(body),"message")//gjson包用来抓取信息
    fmt.Println(result)
    fmt.Println(gjson.Get(string(body),"description"))

    url1 := "https://www.douban.com/people/JacksieCheung/"
    client:=&http.Client{}//创建一个client类型(客户端)变量
    req2,err := http.NewRequest("POST",url1,payload)//通过NewRequest函数发送请求，第一个参数是发送方法，第二个是网址，第三个不知道是什么（暂时）返回响应报文要确定是否为客户>端登陆，存储在req里面。这个payload其实没什么用
    if err != nil {
        fmt.Println("获取地址错误")
        panic(err)
    }
//这里是关键的关键!!!
    var cookie =res.Cookies()
    for _,v := range cookie {
        req2.AddCookie(v)
    }
    //req2.AddCookie(cookie)//在req报文头部加上cookie（用来确定客户端的一段信息）req应该是一种变量类型，Header是方法。
    req2.Header.Add("User-Agent",UserAgent)//在头部加上Agent，用来身份验证。
    resp,err := client.Do(req2)//通过Do方法从客户端发送请求，用resp存储响应报文
    if err != nil {
        fmt.Println("登陆错误")
        panic(err)
    }

    //client.Do发送请求后，后面的步骤就一样了。通过readall读取，然后编辑
    resp_byte,err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    respHtml := string(resp_byte)
    //fmt.Println(respHtml)
    dom,err := goquery.NewDocumentFromReader(strings.NewReader(respHtml))
    if err != nil {
        panic(err)
    }
    dom.Find("a").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
