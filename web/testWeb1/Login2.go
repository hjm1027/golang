//实现数据持久化
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strings"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"time"
	//"encoding/csv"
)

//建立一个map来存储用户信息
var Users = make(map[string]string)
var RealName string

//这是一个程序初始化处理，通过这个把本地的数据先导入map，再运行web
func init() {
	data, err := ioutil.ReadFile("/home/jacksie/desk/golang/web/testWeb1/dataBase1.csv")
	if err != nil {
		fmt.Println(err)
	}
	//把json反序列化变成map
	err = json.Unmarshal([]byte(string(data)), &Users)
}

//登陆板块
func Login(w http.ResponseWriter, r *http.Request) {
	//通过访问web的方式差别进行处理
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login1.html")
		t.Execute(w, nil)
	} else {
		var username, password string
		r.ParseForm() //这是对传入的表单格式进行解析
		for _, v := range r.Form["username"] {
			username = v
			RealName = v //设定一个全局变量便于存储map
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
			//简易cookie的实现
			tNow := time.Now()
			cookie := http.Cookie{Name: "name", Value: RealName, Expires: tNow.AddDate(1, 0, 0)}
			http.SetCookie(w, &cookie)
			//在终端显示信息（用来看看处理到哪了，显得高级一点）
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			//通过html跳转页面
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

//注册板块
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
			fmt.Fprintln(w, "Two passwords should not be the same")
			return
		}

		Users[newuser] = password
		byteUsers, err := json.Marshal(Users) //这里是把map变成json，以存入文件
		if err != nil {
			fmt.Println(err)
		}
		//采用先删除再新建的方法
		del := os.Remove("/home/jacksie/desk/golang/web/testWeb1/dataBase1.csv")
		if del != nil {
			fmt.Println(del)
		}
		file, err := os.Create("/home/jacksie/desk/golang/web/testWeb1/dataBase1.csv")
		if err != nil {
			fmt.Println(err)
		}
		file.Write(byteUsers)
		file.Close()

		t, _ := template.ParseFiles("register2.html")
		t.Execute(w, nil)
		fmt.Fprintln(w, "regist successfully")
	}
}

//修改密码板块
func changePassword(w http.ResponseWriter, r *http.Request) {
	//简易判断有没有cookie
	cookiename, err := r.Cookie("name")
	fmt.Println(cookiename, err)
	if err != nil {
		fmt.Println("No cookie", err)
		http.Redirect(w, r, "http://localhost:9090/login", http.StatusFound) //使用函数跳转
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

		byteUsers, err := json.Marshal(Users)
		if err != nil {
			fmt.Println(err)
		}
		del := os.Remove("/home/jacksie/desk/golang/web/testWeb1/dataBase1.csv")
		if del != nil {
			fmt.Println(del)
		}
		file, err := os.Create("/home/jacksie/desk/golang/web/testWeb1/dataBase1.csv")
		if err != nil {
			fmt.Println(err)
		}
		file.Write(byteUsers)
		file.Close()

		t, _ := template.ParseFiles("changePassword2.html")
		t.Execute(w, nil)
		fmt.Fprintln(w, "Password changed successfully")
	}
}

//主页板块
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
