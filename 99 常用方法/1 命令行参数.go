package main

import (
	"fmt"
	"os"
)

func main() {

	list := os.Args
	fmt.Println(list) // 返回: 文件名 参数.....
	fmt.Println(len(list))
}
