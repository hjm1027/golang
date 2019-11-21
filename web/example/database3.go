//测试删除table和删除用户的行
package main

import (
    "fmt"
    "database/sql"
    "log"
    "time"

    _"github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "JacksieCheung:15811852133@(127.0.0.1:3306)/JacksieCheung?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    {//删除行，加一个for循环可以实现多次删除
        //for n:=3;n<=5;n++{
        _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
        if err != nil {
            log.Fatal(err)
        }
    //}
    }

    { // Query all users
        type user struct {
            id        int
            username  string
            password  string
            createdAt time.Time
        }

        rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var users []user
        for rows.Next() {
            var u user

            err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
            if err != nil {
                log.Fatal(err)
            }
            users = append(users, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%#v", users)
    }
}
