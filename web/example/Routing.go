//example from the GO Web Example
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//这个是初始化新的路由，返回一个r的类型，可以用r.HandleFunc设置路由
	r := mux.NewRouter()

	//HandleFunc用来注册路由
	//第一个参数可以用正则表达式来匹配
	//这里的参数是带变量的参数，title和page部分可以任意
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		//Vars函数返回了一个map，用来存储变量，变量再浏览器的url中获取，并且存在了对应的map里面。这里如果url输入muxistudio，那么map["title"]="muxistudio"
		title := vars["title"]
		page := vars["page"]
		fmt.Println(title, page) //这里的title和page是网址输入的变量值。

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	//这里第二个参数一定是r
	http.ListenAndServe(":9090", r)
}

/*通过调用Method方法还可以实现限制http访问方式
对应不同的handler可以更好的处理网页路由，真正实现不同路由对应不同处理，而不再需要通过if去区别http访问方法。
  这种限定访问方式的，适合对一个网页同时进行get，post，put，delete等多种请求，如果只是两个，用if其实更卡，而且如果handler很简单，还是if更好，甚至switch。但是其实当handler相当复杂，method确实更好。
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

  通过host方法可以将请求处理程序（handler）限制再特定的主机名或子域。
其实把localhost打全，也可以。
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

  将请求处理程序限制为http/https。schemes(方案）
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

  将请求处理程序限制为特定的路径前缀
  pathPrefix是去掉books后面的url
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
*/
