//这里实现登陆和主界面
//直接服务端里搞一个map用来存储
package main

import (
	"fmt"
	"log"
	"net/http"
	//"strings"
	"html/template"
	"time"
)

var Users = make(map[string]string)
var RealName string

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login1.html")
		t.Execute(w, nil)
	} else {
		var username, password string
		r.ParseForm()
		for _, v := range r.Form["username"] {
			username = v
			RealName = v
		}
		for _, v := range r.Form["password"] {
			password = v
		}
		if _, ok := Users[username]; ok {
			if password != Users[username] {
				t, _ := template.ParseFiles("login1.html")
				t.Execute(w, nil)
				fmt.Fprintln(w, "Wrong password")
				return
			}
			tNow := time.Now()
			cookie := http.Cookie{Name: "name", Value: RealName, Expires: tNow.AddDate(1, 0, 0)}
			http.SetCookie(w, &cookie)
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			t, _ := template.ParseFiles("login2.html")
			t.Execute(w, nil)
			fmt.Fprintln(w, "login successfully")
		} else {
			t, _ := template.ParseFiles("login1.html")
			t.Execute(w, nil)
			fmt.Fprintln(w, "Login fail")
			return
		}
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("register1.html")
		t.Execute(w, nil)
	} else {
		var newuser, password, password2 string
		r.ParseForm()
		for _, v := range r.Form["username"] {
			newuser = v
		}
		if _, ok := Users[newuser]; ok {
			t, _ := template.ParseFiles("register1.html")
			t.Execute(w, nil)
			fmt.Fprintln(w, "This username has already exited")
			return
		}
		for _, v := range r.Form["password"] {
			password = v
		}
		for _, v := range r.Form["password2"] {
			password2 = v
		}
		if password != password2 {
			t, _ := template.ParseFiles("register1.html")
			t.Execute(w, nil)
			fmt.Fprintln(w, "Two passwords are not the same")
			return
		}
		Users[newuser] = password
		t, _ := template.ParseFiles("register2.html")
		t.Execute(w, nil)
		fmt.Fprintln(w, "regist successfully")
	}
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	cookiename, err := r.Cookie("name")
	fmt.Println(cookiename, err)
	if err != nil {
		fmt.Println("No cookie", err)
		http.Redirect(w, r, "http://localhost:9090/login", http.StatusFound)
		return
	}
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("changePassword1.html")
		t.Execute(w, nil)
	} else {
		var password1, password2 string
		r.ParseForm()
		for _, v := range r.Form["password1"] {
			password1 = v
		}
		for _, v := range r.Form["password2"] {
			password2 = v
		}
		if password1 == password2 {
			t, _ := template.ParseFiles("changePassword1.html")
			t.Execute(w, nil)
			fmt.Fprintln(w, "Two passwords are the same")
			return
		}
		Users[RealName] = password2
		t, _ := template.ParseFiles("changePassword2.html")
		t.Execute(w, nil)
		fmt.Fprintln(w, "Password changed successfully")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	cookiename, err := r.Cookie("name")
	fmt.Println(cookiename, err)
	if err != nil {
		fmt.Println("No cookie", err)
		http.Redirect(w, r, "http://localhost:9090/login", http.StatusFound)
		return
	}
	fmt.Println("method:", r.Method)
	t, _ := template.ParseFiles("home1.html")
	t.Execute(w, nil)
	fmt.Fprintln(w, "Welcome home! "+RealName)
}

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/home", home)
	http.HandleFunc("/changePassword", changePassword)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
