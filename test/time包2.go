package main

import (
		"fmt"
		"time"//调用time包
)

func main() {
		fmt.Println("When's Saturday?")
		today :=time.Now().Weekday()//time包的now可以获取现在的时间
		fmt.Printf("(%T)",today)	//这里的weekday并不是指工作日，而是指周几，返回类型为time.Weekday型
		switch time.Saturday {//判断时是 time.Saturday==today+1?
		case today+0 :fmt.Println("Today.")
		case today+1 :fmt.Println("Tomorrow.")
		case today+2 :fmt.Println("In two days.")
		default:fmt.Println("Too far away")
		}
}
/*time.weekday类型可以做运算，强制转int得到偏差数，
默认是sunday开始到saturday算0123456，所以这里today+1
是可以运算的，如果是sunday，就要减了。如果是其他就要
判断哪个大再运算*/
