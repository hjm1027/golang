//测试服务器搭建
package main

import (
    "net/http"
    "fmt"
    //"strings"
    "log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){//第一个是响应对象，第二个参数是请求对象
   // r.ParseForm() //解析参数，默认是不会解析的
    //下面一堆这种输出都是可以没有的
    /*fmt.Println(r.Form)
    fmt.Println("path",r.URL.Path)
    fmt.Println("scheme",r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k,v:= range r.Form{
        fmt.Println("key:",k)
        fmt.Println("val:",strings.Join(v,""))
    }*/
    //输出客户端，w是第一个参数
    fmt.Fprint(w,"Hello World!") //w的格式化写入（就是把内容变成前面一个参数的格式）是输出到客户端的
}

func main(){
    http.HandleFunc("/",sayhelloName) //设置访问的路由
    //函数listenAndServe：建立服务器，第一个参数是地址，第二个是处理请求的接口，类型为Hander
    //该函数一直运行除非发生错误，这个程序也会，除非手动关闭或者错误
    //通过ListenAndServe来启动服务
    err := http.ListenAndServe(":9090",nil) //设置监听的端口 localhost:9090
    if err!= nil{
        log.Fatal("ListenAndServe: ",err)
    }
}
