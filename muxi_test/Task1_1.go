//更加完善的版本
package main

import "fmt"

func strStr(haystack string, needle string) int {
        var n int
    S1,S2:=[]byte(haystack),[]byte(needle)
    if len(S1)==0&&len(S2)==0 {
        return 0
    }
    if len(S1)==0&&len(S2)!=0 {
        return -1
    }
    if len(S1)!=0&&len(S2)==0 {
        return 0
    }
    if len(S1)<len(S2) {
        return -1
    }
    for i:=0;i<len(S1);i++ {
        if S1[i]==S2[0] {
            k:=i+1
            for n=1;n<len(S2);n,k = n+1,k+1{
                if k==len(S1) {
                    break
                }
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
    fmt.Println(strStr(str1,str2))
}
