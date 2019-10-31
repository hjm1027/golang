package main

import (
	"fmt"
	"runtime"//导入runtime包
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {//runtime包的GOOS会返回当前操作系统名字
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

