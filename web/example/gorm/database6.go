//测试gorm删除表
package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
)

type Users struct {
    username string `gorm:"type:varchar(128);not null;index:username_idx"`
    password string `gorm:"type:varchar(128);not null;"`
}

func main() {
    db,err := gorm.Open("mysql","JacksieCheung:15811852133@/JacksieCheung?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    db.CreateTable(&Users{})
    fmt.Println("create successfully")
    db.DropTable("users")
}
