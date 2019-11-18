//登陆界面实例
//网页跳转是前段实现的！！！
package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)
//这个过程应该是两个函数都执行了，只不过第一个函数在登陆之前什么都没有执行出来
func sayHelloName(w http.ResponseWriter, r *http.Request) {
    // 解析url传递的参数
    r.ParseForm()//这里的r其实是一个表单数据
    for k, v := range r.Form {//这个表单传进来应该是一个map
        fmt.Println("key:", k)
        // join() 方法用于把数组中的所有元素放入一个字符串。
        // 元素是通过指定的分隔符进行分隔的
        //这里的值是字符串切片[]string类型，用join转化为string
        fmt.Println("val:",strings.Join(v,""))//这里通过拆分再合并把值变成string
    }
    // 输出到客户端
    ////r.form找到request请求里面的表单部分，这里访问的是键
    name :=r.Form["username"]
    pass :=r.Form["password"]
    for i,v :=range name{
        fmt.Println(i)//这里是输出在终端
        fmt.Fprintf(w,"用户名:%v\n",v)//带F的是写入w，也就是写入响应报文的boady
    }
    for k,n :=range pass{
        fmt.Println(k)
        fmt.Fprintf(w,"密码:%v\n",n)
    }
}
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    //通过这一段可以实现登陆才能进网页
    if r.Method == "GET" {//这里用请求方法来分类，只是单纯获取界面，就展示html
        t, _ := template.ParseFiles("testLogin1.html")
        // func (t *Template) Execute(wr io.Writer, data interface{}) error {
        t.Execute(w, nil)//这个是执行html文件，换句话说就是跳转到登陆页面。
    } else {//输入密码会变成POST请求或者别的，这个时候就是另外一种处理。
        r.ParseForm()//登陆之后又会是get
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}
func main() {
    http.HandleFunc("/hello", sayHelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe("localhost:9090", nil)
    if err != nil {
        log.Fatal("ListenAndserve:", err)
    }
}
