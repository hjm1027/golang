//测试查询返回的err
package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
)

type Users struct {
    //这里的成员都要大写，不然表建不出来。
    //这里not null意思是这列必须不为空，还可以有AUTO_INCREMENT 表示该属性自增
    //自定义列名：comlumn:<列名>
    Username string `gorm:"type:varchar(128);not null;"`
    Password string `gorm:"type:varchar(128);not null;"`
    Email string `gorm:"type:varchar(128);not null;"`
}


func main() {
    db,err := gorm.Open("mysql","JacksieCheung:15811852133@/JacksieCheung?charset=utf8&parseTime=True&loc-Local")//true才是对的，不是ture！！！
    if err != nil {
        panic(err)
    }
    defer db.Close()
    err2 := db.CreateTable(&Users{}).Error
    if err2 != nil {
        panic(err2)
    }
    /*user1 := &Users {
        Username :"JacksieCheung",
        Password:"15811852133",
        Email:"1799902904@qq.com",
    }
    if err:=db.Create(user1).Error;err!=nil {
        panic(err)
    }*/
    //如果找不到，会返回true，找到了会返回false
    test := db.Where("username=?","JacksieCheung").Find(&Users{}).RecordNotFound()
    fmt.Println(test)
}
