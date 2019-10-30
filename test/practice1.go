//用循环和分支写函数实现牛顿开方函数
package main

import "fmt"

func sqr(x float64) {
		var i int = 0
		var z float64 = 1
		var min float64
		min = z-(z*z-x)/(2*z)
		for i=0;i<10;i++{
				z=z-(z*z-x)/(2*z)
				if i==0 {
						continue
				}
				if z<min {
						min=z
						continue
				}
				if z==min {
						break
				}
		}
		fmt.Println("一共计算的次数:",i)//只有printf格式化输出才能用%f的形式
		fmt.Println("结果:",min)
}

func main() {
		var x float64
		fmt.Println("请输入一个数用来求它的开方")
		fmt.Scanf("%f",&x)
		sqr(x)
}
