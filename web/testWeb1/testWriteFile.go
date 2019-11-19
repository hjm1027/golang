//测试写入map
package main

import (
    "fmt"
    "encoding/json"
    "os"
    "log"
    "io"
    "encoding/csv"
)

func main() {
    users := map[string]string{
        "JacksieCheung":"123456789",
        "Lucy":"256789",
        "Kiki":"85674123",
    }
    str,err := json.Marshal(users)
    if err != nil {
        fmt.Println(err)
    }

    var strs []string
    for i:=0;i<len(str);i++{
        strs = append(strs,string(str[i]))
    }

    newFileName:="/home/jacksie/desk/golang/web/testWeb1/test.csv"

    nfs, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        log.Fatalf("can not create file, err is %+v", err)
    }
    defer nfs.Close()
    nfs.Seek(0, io.SeekEnd)

    w := csv.NewWriter(nfs)
    //写入csv
    err = w.Write(strs)
    if err != nil {
        log.Fatalf("can not write, err is %+v", err)
    }
    //这里必须刷新，才能将数据写入文件。
    w.Flush()
}
