//函数的闭包测试
package main

import "fmt"

func add(x int,y int) func(int) int {
		x = 50
		y = 60
		return func(z int) int{
				return x+y+z
		}
}

func main() {
		var x,y,z int
		fmt.Println(add(x,y)(z))//可以通过圆括号后面圆括号这种形式导入嵌入函数的参数。
}

/*func add() func(int) int {
		i:=100
		return func(x int) int {
				return x+i
		}
}

func main() {
		A := add()
		fmt.Println(A(5))
}*/
//另外函数名是一个指针，指向函数的地址。然后当声明函数变量给它赋值例如：A:=add(x,y)这种还把参数给确定了，这样A就指向了这个函数的里层的第一层函数（这里也是最后一层函数）。
