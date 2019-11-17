//测试声明类型变量再设置访问路由
package main

import (
    "log"
    "net/http"
)

type MyHttpHandler struct{}

func (handler MyHttpHandler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
    switch r.URL.Path {
        case "/a":
            w.Write([]byte("<html><body>Hello,world!</body></html>"))
        case "/b":
            w.Write([]byte("<html><body>Goodbye,world!</body></html>"))
        }
}

func main() {
    a:=MyHttpHandler{}
    b:=MyHttpHandler{}
    http.HandleFunc("/a",a.ServeHTTP)
    http.HandleFunc("/b",b.ServeHTTP)
    log.Fatal(http.ListenAndServe("localhost:9090",nil))
}
