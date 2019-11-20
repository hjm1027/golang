//测试连接mysql
package main

import (
    "database/sql"
     _ "go-sql-driver/mysql"
     "log"
 )

func main() {
// Configure the database connection (always check errors)
    db, err := sql.Open("mysql", "JacksieCheung:15811852133@(127.0.0.1:3306)/JacksieCheung?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
    if err := db.Ping();err != nil {
        log.Fatal(err)
    }
}
