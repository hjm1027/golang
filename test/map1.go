//map 对键值对进行操作，是存储容器
//map是一种无序的键值对集合，map最重要一点是通过key来快速检索数据，key类似于索引，指向数据的值。map是一种集合，所以我们可以像迭代数组和切片一样去迭代它。但是map是无序的，我们无法决定它的返回顺序，这是因为map是使用哈希表（hash）来实现的。
package main

import "fmt"

func main() {
		var CapitalMap map[string]string
//map变量的声明格式是map[]__,[]里面是键的类型，后面跟着的是值的类型（甚至可以是结构体类型）。值和建是对应关系，键是值的索引
//还有另外一种声明，var CapitalMap=make(map[string]string),不能用new建
		CapitalMap = make(map[string]string)
		//map类型赋值的实现(个数不受限制)
		//有两种赋值方式，第一种是一个一个赋，第二种是一次性
		CapitalMap["China"]="Beijing"
		CapitalMap["Japan"]="Tokey"
		//如果是结构体，那么要用大括号 m[]={  },如果是初始化一次性赋值，冒号后面加上大括号 “ ”：{ } 
		var CapitalMap2 =map[string]string{"France":"Paris","Italy":"Rome"}//第二种只能用作初始化
		//可以通过遍历来实现输出
		for country :=range CapitalMap {//这里的country是索引，把值去掉了
				fmt.Println(country,"首都是",CapitalMap[country])
		}
		for country :=range CapitalMap2 {
				fmt.Println(country,"首都是",CapitalMap2[country])
		}
		//删除元素的实现 delete（）函数
		delete(CapitalMap2,"France")
		for country :=range CapitalMap2 {
				fmt.Println(country,"首都是",CapitalMap2[country])
		}
}
//值为nil的map是空且不能赋值

