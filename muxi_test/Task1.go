//找字符串中子字符串位置，返回位置
//不用内置函数做法
package main

import "fmt"

func FindSubstring(str1,str2 string) int {
    var n int
    S1,S2:=[]byte(str1),[]byte(str2)
    for i:=0;i<len(S1);i++ {
        if S1[i]==S2[0] {//如果遇到字母相等，单独抽出来循环
            k:=i+1
            for n=1;n<len(S2);n,k = n+1,k+1{//++是语句，这里需要表达式。单独一个i++可以，但是多个不行。多个需要写成如代码所示。
                if S1[k]!=S2[n]{
                    break
                }
            }
            if n==len(S2) {
                return i
            }
            n=0
        }
    }
    return -1
}

func main() {
    var str1,str2 string
    fmt.Scanf("%s%s",&str1,&str2)
    fmt.Println(FindSubstring(str1,str2))
}
