//用来登陆的片段
package login

import (
    "net/http"
)

type Data struct{
    name string
    password string
}

var User Data

func Login(w http.ResponseWriter,r *http.Request) {
    r.ParseForm{}
    User.Name := r.FormValue("Username")//参数和返回值都是string
    User.Password := r.FormValue("Password")
    if _,ok := dataBase.Users[User.Name],ok {
        fmt.Println("
