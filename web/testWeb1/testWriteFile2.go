//测试写入map
package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"log"
	//"io"
	//"encoding/csv"
)

func main() {
	users := map[string]string{
		"JacksieCheung": "123456789",
		"Lucy":          "256789",
		"Kiki":          "85674123",
	}
	str, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	del := os.Remove("/home/jacksie/desk/golang/web/testWeb1/test.csv")
	if del != nil {
		fmt.Println(del)
	}

	file, err := os.Create("/home/jacksie/desk/golang/web/testWeb1/test.csv")
	if err != nil {
		fmt.Println(err)
	}
	//写入byte的slice数据
	file.Write(str)
	file.Close()
}
