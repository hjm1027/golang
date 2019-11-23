//用gorm的数据库处理测试
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"//需要导入数据库驱动（driver）不然连不上数据库
    //_"github.com/jinzhu/gorm/dialects/sqlite"
)

/*type Product struct {
    gorm.Model
    Code string
    Price uint
}*/

var db *gorm.DB

func main() {
    var err error
    db,err = gorm.Open("mysql","JacksieCheung:15811852133@/JacksieCheung?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("err")
    }
}
