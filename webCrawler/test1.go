//爬取豆瓣前250电影
package main

import (
    "fmt"
    "strconv"

    "github.com/PuerkitoBio/goquery"

)

func main() {
    for page := 0;page<=9;page++{
        ur1 := "http://movie.douban.com/top250?start="+strconv.Itoa(page*25)+"&filter="
        opening,_:=goquery.NewDocument(ur1)
        ele:=opening.Find("div.hd")
        ele.Each(func(index int,content *goquery.Selection) {
            filename := content.Find("a").Find("span.title").Eq(0).Text()
            fmt.Println(index+1+page*25,filename)
        })
    }
}
