//测试爬虫
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    URL1:="http://search.zongheng.com/s?keyword=火影&pageNo=1&sort="
    rp,err:=http.Get(URL1)
    if err != nil {
        panic(err)
    }
    body,err:=ioutil.ReadAll(rp.Body)
    content := string(body)//把body转换成了字符串类型,如果不转换，那就是网页源代码的ascii码。
    defer rp.Body.Close()
    fmt.Println(content)
}
