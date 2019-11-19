//测试csv文件读写
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

func main() {
    b,err := ioutil.ReadFile("/home/jacksie/desk/golang/web/testWeb1/test.csv")
    if err!=nil {
        fmt.Println(err)
    }

    list := map[string]string{}
    err=json.Unmarshal([]byte(string(b)),&list)
    if err != nil {
        panic(err)
    }
    fmt.Println(list)
}
