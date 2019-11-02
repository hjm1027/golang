//测试string类型改变的一个示例
package main

import "fmt"

func main() {
		strings := "my name is Jacksie"
		char :=[]byte(strings)//可以通过这样初始化直接把字符串转化为字符数组
		char[0]='M'
		strings =string(char)//修改字符串要先把改的那一段切出来或者整个拿出来存到字符数组里面。然后再用原来字符串存修改后的数组。对于byte类型。它应该有两个值。一个是字符，一个是ASCII表对应int。Println输出时候默认是int类型。要输出字符就要用printf格式化输出。在go的ASCII运算如 '0'+1等都是把字符转成int运算。而c刚好相反。所以要加上强制类型转换，前面加个byte（）
		fmt.Println(strings)
}
