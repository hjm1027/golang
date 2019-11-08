//最终版（go就应该这样写）
package main

import "fmt"

func strStr(haystack string, needle string) int {
    if len(needle)==0 {
        return 0
    }
    if len(haystack)<len(needle) {
        return -1
    }
    for i:= 0;i<len(haystack);i++{
        //直接用haystack[:]切出来一块比较就可以了。但这并不是切片，只是访问而已
        //切片只能和nil比较，slice之间不能比较。所以这个不是切片。
        if len(needle)+i<=len(haystack)&&haystack[i:i+len(needle)]==needle {
            return i
        }
    }
    return -1
}

func main() {
    var str1,str2 string
    fmt.Scanf("%s%s",&str1,&str2)
    fmt.Println(strStr(str1,str2))
}
