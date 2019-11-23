//测试gorm创建table
package main

import (
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
    //"time"
)

type Users struct {
    Username string `gorm:"type:varchar(128);not null;index:username_idx"`
    Password string `gorm:"type:varchar(128);not null;"`
    //CreatedAt time.Time
}

func main() {
    db,err := gorm.Open("mysql","JacksieCheung:15811852133@/JacksieCheung?charset=utf8&parseTime=True&loc=Local")
    defer db.Close()
    if err != nil {
        panic(err)
    }
    /*user := &Users{
        ID:2019214228,
        username:JacksieCheung,
        password:15811852133,
    }*/
    /*if !db.HasTable(&Users{}) {//可以前面跟着set，设置属性
        if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Users{}).Error; err != nil {
            panic(err)
        }
    }*/
    db.CreateTable(&Users{})//这个是没有返回错误的
}

