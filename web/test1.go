//测试本地服务端实现
package main

import (
    "flag"
    //"fmt"
    "net/http"
)

func main() {
    host := flag.String("host","127.0.0.1","listen host")//127.0.0.1就是本地的地址
    port := flag.String("port","9090","listen port")//设置接口

    http.HandleFunc("/hello",Hello)//这个函数是设置域名，第一个参数设置了hello域名，访问只要在原来地址后面加上域名就可以了，第二个参数是处理函数（handler）

    err := http.ListenAndServe(*host+":"+*port,nil)//设置监听端口，有两种方式，第一种是第一个参数像这样完整的，还有一种可以直接":9090"设置。

    if err != nil {
        panic(err)
    }
}

func Hello(w http.ResponseWriter,req *http.Request) {
    w.Write([]byte("Hello World"))//write方法写入
}
