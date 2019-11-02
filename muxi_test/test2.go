//整数转换为字符串
package main

import "fmt"

func main() {
		var n int
		count :=0
		fmt.Scanf("%d",&n)
		var char = make ([]byte,0)
		for n!=0 {
				char =append(char,byte(n%10+'0'))//通过ASCII表的转化是存在的，但要加一个类型转换。因为go默认转化为int类型。
				n/=10
				count++
		}
		for i:=count-1;i>=0;i--{
			fmt.Printf("%c",char[i])
	}
}
