//搭建服务器一般流程
package main

import (
    "log"
    "net/http"
)

type MyHttpHandler struct{}//也可以是map类型

//实现ServeHTTP接口
//这里是声明了一个叫做MyHttpHandler的结构体类型，然后这里是ServeHTTP方法声明
//后面一定要跟这个方法
func (handler MyHttpHandler)ServeHTTP(w http.ResponseWriter,r*http.Request) {
    //调用write方法把html写入响应流。当然也可以用fmt包的fprint
    //可以直接在实现接口时写一个for循环来找r（辨别请求地址）
    w.Write([]byte("<html><body>Hello,world!</body></html>"))
}

func main() {
    http.HandleFunc("/",MyHttpHandler{}.ServeHTTP)//这个函数的第二个参数接收的是一个方法或函数。这是为了设置访问路由，就是访问的路径，在后面加域名。
    //ListenAndServe(address string,h Handler) error 这是启用服务器的函数
    //h是处理请求的接口，类型为Handler adress是服务器的地址
    //这里MyHttpHandler是自定义声明，要使用一般要声明一个这个类型的变量，但是这里也可以没有声明变量而参数直接代入这个类型
    log.Fatal(http.ListenAndServe("localhost:9090",nil))
}
