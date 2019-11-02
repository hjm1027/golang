//整数转换为字符串
package main

import "fmt"

func main() {
        var n,k,count int
        fmt.Scanf("%d",&n)
        var char = make ([]byte,0)
        for n!=0 {
                char =append(char,byte(n%10+'0'))
                n/=10
                count++
        }
		var chars =make([]byte,count)//[]里面不能是变量
        for i:=count-1;i>=0;i--{
				chars[k]=char[i]
			k++
    }
	fmt.Println(string(chars[:]))//字符数组不能直接转换为string，但是切片可以，而数组又可以直接转成切片。
}

