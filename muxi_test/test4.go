//今天一年的第几天 YYYY-MM-DD
package main

import "fmt"

func main() {
	var days =[]int{31,28,31,30,31,30,31,31,30,31,30,31}
	var date =make([]byte,10)
	for i:=0;i<10;i++{
	fmt.Scanf("%c",&date[i])
	}
	var year,month,day int
	for i:=0;i<4;i++{
			year=year*10+int(date[i]-'0')
	}
	for i:=5;i<7;i++{
			month=month*10+int(date[i]-'0')
	}
	for i:=8;i<10;i++{
			day=day*10+int(date[i]-'0')
	}

	if year%4==0{
			days[1]=29
	}

	var sum int
	for i:=0;i<month-1;i++{
			sum+=days[i]
	}
	fmt.Println(sum+day)
}
