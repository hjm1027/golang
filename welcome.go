//只是写一个小程序玩玩
package main

import(
		"fmt"
)
func main(){
		var array string
		fmt.Println("Please input yout name")
		fmt.Scanf("%s",&array)
		if(array=="JacksieCheung"){
				fmt.Println("welcome back,my dear friend")
		}	else{
				fmt.Println("Who are you?\nIt's my first time to see you.")
		}
}
