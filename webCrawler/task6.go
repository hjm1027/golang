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
//声明一个url包里的values类型，并用方法Set设置body的数据
    data := url.Values{}
    data.Set("name",name)
    data.Set("password",string(password))
    data.Set("remember","false")
    data.Set("ck","")
    data.Set("ticket","")

    payload := strings.NewReader(data.Encode())//把用户信息放进payload里

    req1,err := http.NewRequest("POST",requestUrl1,payload)//建立一个新的请求
    if err != nil {
        panic(err)
    }
//加上请求头
    req1.Header.Add("Accept","application/json")
    req1.Header.Add("Content-Type","application/x-www-form-urlencoded")
    req1.Header.Add("Origin","https://accounts.douban.com")
    req1.Header.Add("Referer","https://accounts.douban.com/passport/login_popup?login_source=anony")
    req1.Header.Add("Sec-Fetch-mode","cors")
    req1.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
    req1.Header.Add("X-Requested-with","XMLHttpRequest")
//客户端发送请求
    res1,err := http.DefaultClient.Do(req1)
    if err != nil {
        panic(err)
    }
    defer res1.Body.Close()
//读取响应报文的body部分，目的是读取响应信息（读取cookie在下面）
    body,err := ioutil.ReadAll(res1.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(res1.Status)//这里是响应状态的说明
    fmt.Println(gjson.Get(string(body),"message"))
    fmt.Println(gjson.Get(string(body),"description"))

//现在开始模拟登陆已经完成，响应报文已经存在了res里面，接下来开始拿着res里的cookie爬取目标网页 

    requestUrl2 := "https://www.douban.com/people/JacksieCheung/"
    client := &http.Client{}//声明一个客户端（client）类型的变量
    req2,err := http.NewRequest("POST",requestUrl2,nil)//建立一个新的请求
    if err != nil {
        panic(err)
    }
//把cookie从响应报文里面抓出来，通过AddCookie方法加在请求2里
    var cookie =res1.Cookies()
    for _,v := range cookie {
        req2.AddCookie(v)
    }
//设置代理（userAgent）这个是不变的，直接从浏览器里面复制就可以了
    userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36"
    req2.Header.Add("User-Agent",userAgent)//在头部加上代理
    res2,err := client.Do(req2)//客户端发送请求
    if err != nil {
        panic(err)
    }
//后面的操作就和普通爬虫一样了
    res2_byte,err := ioutil.ReadAll(res2.Body)
    defer res2.Body.Close()
    respHtml := string(res2_byte)

    dom,err := goquery.NewDocumentFromReader(strings.NewReader(respHtml))
    if err != nil {
        panic(err)
    }
    dom.Find("a").Each(func(i int,s *goquery.Selection) {
        fmt.Println(s.Text())
    })
}
