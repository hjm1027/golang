//随便从百度扒东西
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    rp,err:=http.Get("http://www.baidu.com")
    if err!=nil {
        panic(err)
    }
    body,err:=ioutil.ReadAll(rp.Body)
    defer rp.Body.Close()//引用要大写
    fmt.Printf("%s",body)
}
