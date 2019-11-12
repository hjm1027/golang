//测试网页搜索
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    //"strings"

    //"github.com/PuerkitoBio/goquery"
    "regexp"
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

   // dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        panic(err)
    }
    reg := regexp.MustCompile(` </span><span title="up主" class="so-icon"><i class="icon-uper"></i><a href="//space.bilibili.com/32660634?from=search" target="_blank" class="up-name">bilibili星访问</a></span></div></div></li><li class="video-item matrix"><a href="//www.bilibili.com/video/av73262911?from=search" title="【官方】《天气之子》主题曲《爱能做到的还有什么》MV（演唱：RADWIMPS）" target="_blank" class="img-anchor"><div class="img"><div class="lazy-img"><img alt="" src=""></div><span class="so-imgTag_rb">02:49</span><div class="watch-later-trigger watch-later"></div><span class="mask-video"></span></div><!----></a><div class="info"><div class="headline clearfix"><!----><!----><span class="type hide">MV</span><a title="【官方】《天气之子》主题曲《爱能做到的还有什么》MV（演唱：RADWIMPS）" href="//www.bilibili.com/video/av73262911?from=search" target="_blank" class="title">【官方】《<em class="keyword">(.*?)</a></div><div class="des hide">`)
    result := reg.FindAllStringSubmatch(dom,-1)
    fmt.Println(result)
}
