//矩阵
package main

import "fmt"

func main() {
		var i,j int
		fmt.Scanf("%d%d",&i,&j)//不知道怎么限定行和宽再建立数组
		var array ={[][]float64,i*j}

		for a:=0;a<i;a++{
				for b:=0;b<j;b++{
						fmt.Scanf("%d",&array[i][j])
				}
		}
		fmt.Println(array)
}
