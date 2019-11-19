//test from web eample
//文件浏览器的实现
package main

import (
    "fmt"
    "net/http"
    //"os"
    //"path/filepath"
)

func main() {
    /*http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to my website!")
    })*/

    //p,_ :=filepath.Abs(filepath.Dir(os.Args[0]))
    //fmt.Println(p)
    //http.Handle("/",http.FileServer(http.Dir(p)))
    //http.Handle("/b",http.StripPrefix("/b/",http.FileServer(http.Dir(p))))

    //这里的路径是绝对路径，只能从/home开始，从～开始的那个不行
    fs := http.FileServer(http.Dir("/home/jacksie/desk/golang/web"))
    http.Handle("/web/",http.StripPrefix("/web/",fs))//后面的路径和前面的一致

    http.HandleFunc("/a",route)
    http.ListenAndServe("localhost:9090", nil)
}


func route(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL)
    fmt.Fprintln(w, "welcome")
    r.Body.Close()
}
//几个函数或方法的使用
//dir用来生成文件系统。
//FileServer的参数是一个文件系统，返回一个handler（处理器）类型（是一个接口）
//Handle用来开启一个文件服务器。
//StripPrefix用来删除第一个参数（是个路径）后面的url（这个应该是为了避免访问之后的路径）接收一个handler，返回一个handler。并且和handle函数的参数一致，没有时会报错。
