//用Buffers
//实现一个int数组转化成string
package main

import (
		"fmt"
		"bytes"
)

func IntsToString(values []int) string {//该函数和fmt.Sprintf()函数类似，就是多打了个逗号。
		var buf bytes.Buffer//该类型不需要初始化，零值就是可用变量
		buf.WriteByte('[')//在字符串两边加上'['
		for i,v := range values {//v接收数组的值
				if i>0 {
						buf.WriteString(",")
				}
				fmt.Fprintf(&buf,"%d",v)//通过Fprintf函数转换成字符，输入到buf类型里面
		}
		buf.WriteByte(']')
		return buf.String()
}

func main() {
		var value [3]int
		for i:=range value {
				fmt.Scanf("%d",&value[i])
		}
		fmt.Println(IntsToString(value[:]))
}
