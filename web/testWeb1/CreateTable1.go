//先建立一个简单的数据库
package main

import (
    //"fmt"
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
)

type Users struct {
    Username string `gorm:"type:varchar(128);not null"`
    Password string `gorm:"type:varchar(128);not null"`
}

func main() {
    db,err := gorm.Open("mysql","JacksieCheung:15811852133@/JacksieCheung?charset=utf8&parseTime=True&loc-Local")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    err2 := db.CreateTable(&Users{}).Error
    if err2 != nil {
        panic(err)
    }
}
