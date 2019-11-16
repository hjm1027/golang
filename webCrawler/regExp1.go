//test the regular expression
package main

import (
    "regexp"//regular expression
    "fmt"
)

func main() {
    const text = "My email is hanru723@163.com@haha.com"
    //re,err := regexp.Compile("hanru723@163.com")
    //其实这个不难，这里是表达式的组合。[]里面有小写字母a-z，大写字母A-Z，还有0-9。也就是说@和.都不算进去。（\.用来表示.）这样的话就是三个部分，分别通过@和\.连接而成。所以匹配的时候也是找有@和\.的部分。
    //这里去掉\.也可以，是因为这里直接到了@为界，写上会更加规范。
    re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)

    match := re.FindString(text) //查找一个

    fmt.Println(match)
}
