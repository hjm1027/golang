//测试切片扩张
package main

import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)

   numbers = append(numbers, 0)
   printSlice(numbers)

   numbers = append(numbers, 1)
   printSlice(numbers)

   numbers = append(numbers, 2)
   printSlice(numbers)

   numbers = append(numbers, 3)
   printSlice(numbers)

   numbers = append(numbers, 4)
   printSlice(numbers)

   numbers = append(numbers, 5, 6, 7)
   printSlice(numbers)

   numbers = append(numbers, 8)
   printSlice(numbers)

   numbers = append(numbers, 9, 10, 11)
   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
