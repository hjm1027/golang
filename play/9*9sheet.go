//用go打一个九九乘法表
package main

import "fmt"

func main(){
		for i:=1;i<=9;i++ {
				for j:=1;j<=9;j++ {
						fmt.Printf("%3d",i*j)
				}
				fmt.Printf("\n")
		}
}
