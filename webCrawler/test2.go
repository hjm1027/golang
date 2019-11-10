//爬虫测试
package main

import (
    "fmt"
    "net/http"//net/http包能可以发起和接收http网络请求
    "io/ioutil"//io/ioutil包可以解析成我们能够阅读的源码
)

func get() {
    r,err :=http.Get("https://api.github.com/events")//用http.Get()来获取，就是http包里面的Get函数。函数的参数是网址，也可以用变量先把网址存起来：URL1:="httos://api.github.com/events"这应该是一个字符串类型。
//通过get函数返回了两个值，第一个是返回响应报文，这里存在了r里面（其实是一个像结构体一样的东西，里面有网页源代码。第二个是返回是否请求、响应成功，如果成功返回nil
    if err != nil{
        panic(err)//就是直接退出程序，panic函数会等defer运行完再运行
    }
    defer func() {_=r.Body.Close()}()//这个r接收后有成员body，用close关闭文件。
    body,_:=ioutil.ReadAll(r.Body)//用ioutil包的readall函数读取文件本体。这个r有点像结构体，返回了响应报文然后分成了多个成员，其中body就是网页的源代码。
    fmt.Printf("%s",body)
}

func main() {
    get()
}
