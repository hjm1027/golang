//用密码用户名形式模拟登陆
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    //"github.com/howeyc/gopass"
    //"github.com/tidwall/gjson"
)

func main() {
    requestUrl := "https://accounts.douban.com/j/mobile/login/basic"

  /*  var name string
    fmt.Print("豆瓣账号：")
    _,_=fmt.Scanln(&name)
    fmt.Print("输入密码:")*/
    //password,_ :=gopass.GetPasswdMasked()

    data := url.Values{}
    data.Set("name","15811852133")
    data.Set("password",string(/*password*/15811852133))
    data.Set("remember","false")
    data.Set("ck","")
    data.Set("ticket","")

    payload := strings.NewReader(data.Encode())

    req,err := http.NewRequest("POST",requestUrl,payload)
    if err != nil {
        panic(err)
        return
    }

    req.Header.Add("Accept","application/json")
    req.Header.Add("Accept-Encoding","gzip, deflate, br")
    req.Header.Add("Content-Length","64")
    req.Header.Add("Host","accounts.douban.com")
    req.Header.Add("Content-Type","application/json")
    req.Header.Add("Origin","https://accounts.douban.com")
    req.Header.Add("Referer","https://accounts.douban.com/passport/login_popup?login_source=anony")
    req.Header.Add("Cookie","ll=\"118254\"; bid=i-VPVxbCy9g; apiKey=; __utma=30149280.1677139224.1573619422.1573619422.1573619422.1; __utmb=30149280.17.10.1573619422; __utmc=30149280; __utmz=30149280.1573619422.1.1.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/passport/login_popup; login_start_time=1573620046724; last_login_way=account; ap_v=0,6.0; push_noty_num=0; push_doumail_num=0; __utmv=30149280.20662; __utmt=1")
    req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")
    // req.Header.Add("User-Agent","Moz:illa/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0")
    req.Header.Add("X-Requested-with","XMLHttpRequest")
    //req.Header.Add("Sec-Fetch-Mode","cors")

    res,err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
        return
    }
    defer res.Body.Close()

    body,err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
        return
    }
    fmt.Println(res.Status)

    /*result := gjson.Get(string(body),"message")
    fmt.Println(result)
    fmt.Println(gjson.Get(string(body),"description"))*/

    fmt.Println(string(body))
}
