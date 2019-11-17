//实现用户登陆
//写一个文件专门存储用户名信息
//写一个路由作为登陆界面
package main

import (
    "log"
    "net/http"
)

func main() {

    log.Fatal(http.ListenAndServe("localhost:9090",nil))
}
