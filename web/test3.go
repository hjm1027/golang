//实现读取外部html
package main

import (
    "io"
    "net/http"
    "log"
    "os"
    //"fmt"
)

func web(w http.ResponseWriter,r *http.Request) {
    f,err := os.Open("hello.html")//用os包的函数读取html文件
    if err != nil {
        log.Fatalln(err)
    }
    defer f.Close()
    //这里写入html文件不能用fprint，会变成乱码，可以用io包的copy
    io.Copy(w,f)
}

func main() {
    http.HandleFunc("/",web)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("ListenAndServe: ",err)
    }
}
