//对切片重新切片
package main

import "fmt"

func main() {
		p :=[]int{2,3,5,7,11,13}
		fmt.Println("p==",p)
		fmt.Println("p[1:4]==",p[1:4])
		fmt.Println("p[:3]==",p[:3])//省略前面代表从0开始
		fmt.Println("p[4:]==",p[4:])//省略后面代表一直到len(s)结束
		//数组其实也可以达到这种效果，但是要用for循环中间切出来
}
//切片表达式 s[lo:hi] 表示从lo到hi-1的slice元素（含两端）其实就是左闭右开。
//s[lo:lo]是空的，而s【lo：lo+1]只有一个元素
//这样创建了一个新的slice值指向相同的数组。
