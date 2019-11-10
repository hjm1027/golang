package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var keyword string
	fmt.Scanln(&keyword)

	num := 1

	for i := 1; i <= 10; i++ {//这个循环是循环页码用的,itoa把int转化为string，这里的keyword是URL里面的搜索关键字,这里的iota是网页。合起来就是整个网址
		requestUrl := "http://search.zongheng.com/s?keyword=" + keyword + "&pageNo=" + strconv.Itoa(i) + "&sort="
		rp, err := http.Get(requestUrl)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(rp.Body)
		if err != nil {
			panic(err)
		}
		content := string(body)//还是把响应报文的主体拿出来
		defer rp.Body.Close()//关掉文件

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
//goquery包的NewDocumentFromReader()函数返回一个*Document和error，前者代表一个将要被操作的HTML文档。而error返回的是是否调用成功。newreader应该是一个方法(?)
		if err != nil {
			panic(err)
		}
//dom接收了html文件后，（原来的comment不能操作，而dom可以）
		dom.Find(".search-tab").Each(func(i int, selection *goquery.Selection){
			// fmt.Println(selection.Text())
			selection.Find(".tit").Each(func(i int, title *goquery.Selection) {
				// fmt.Println(title.Text())
				fmt.Printf("%3d   ", num)
				fmt.Println(title.Text())
				num++
			})
		})
	}
}
