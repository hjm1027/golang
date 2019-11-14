//模拟登陆豆瓣爬取个人主页
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
//先进行模拟登陆，目的是获取cookie
    requestUrl1 := "https://accounts.douban.com/j/mobile/login/basic"

    var name string
    fmt.Print("请输入豆瓣账号:")
    _,_=fmt.Scanln(&name)
    fmt.Print("请输入密码:")
    password,_ :=gopass.GetPasswdMasked()

    data := url.Values{}
    data.Set("name",name)
    data.Set("password",string(password))
    data.Set("remember","false")
    data.Set("ck","")
    data.Set("ticket","")

    payload := strings.NewReader(data.Encode())

    req,err := http.NewRequest("POST",requestUrl1,payload)
    if err != nil {
        panic(err)
    }

    req.Header.Add("Accept","application/json")
    req.Header.Add("Content-Type","application/x-www-form-urlencoded")
    req.Header.Add("Origin","https://accounts.douban.com")
    req.Header.Add("Referer","https://accounts.douban.com/passport/login_popup?login_source=anony")
    req.Header.Add("Sec-Fetch-mode","cors")
    req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
    req.Header.Add("X-Requested-with","XMLHttpRequest")

    res,err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    body,err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(res.Status)
    fmt.Println(gjson.Get(string(body),"message"))
    fmt.Println(gjson.Get(string(body),"description"))

//现在开始模拟登陆已经完成，响应报文已经存在了res里面，接下来开始拿着res里的cookie爬取目标网页 

    requestUrl2 := "https://www.douban.com/people/JacksieCheung/"
    client := &http.Client{}
    req2,err := http.NewRequest("POST",requestUrl2,nil)
    if err != nil {
        panic(err)
    }

    var cookie =res.Cookies()
    for _,v := range cookie {
        req2.AddCookie(v)
    }
    userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36"
    req2.Header.Add("User-Agent",userAgent)
    resp,err := client.Do(req2)
    if err != nil {
        panic(err)
    }

    resp_byte,err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    respHtml := string(resp_byte)

    dom,err := goquery.NewDocumentFromReader(strings.NewReader(respHtml))
    if err != nil {
        panic(err)
    }
    dom.Find("a").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
