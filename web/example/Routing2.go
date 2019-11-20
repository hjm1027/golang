//测试限制路由方法
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

func Welcomehome(w http.ResponseWriter,r *http.Request) {
    fmt.Println(r.Method)
    fmt.Fprintln(w,"Welcome home")
}

func main() {
    r := mux.NewRouter()
    //如果是别的类型访问，会直接405拒绝。
    r.HandleFunc("/",Welcomehome).Methods("GET")
    //这里第二个参数不能是nil，而应该是r。对应的新路由。
    //如果是nil就会404。因为建立的handler没有与端口产生联系,而不用这个包，直接来就可以。
    err := http.ListenAndServe(":9090",r)
    if err != nil {
        log.Fatal("ListenAndServe:",err)
    }
}
