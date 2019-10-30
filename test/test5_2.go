package main

import (
		"fmt"
		"math/cmplx"
)

var (
		ToBe bool =false
		MaxInt uint64 =1<<64-1
		z complex128 =cmplx.Sqrt(-5+12i)//complex是复数类型
)

func main(){
		i:=1;
		const f="%T(%v)\n"
		//fmt.Printf(f,ToBe,ToBe)
		//fmt.Printf(f,MaxInt,MaxInt)
		//fmt.Printf(f,z,z)
		fmt.Println("%d",i)
		fmt.Println("%v",i)
}
