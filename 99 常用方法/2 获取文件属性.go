package main

import (
	"fmt"
	"os"
)

func main() {

	path := "/Users/weizhenping/Downloads/sd.mp4"
	info, err  := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info.Name())
	fmt.Println(info.Size())
	fmt.Println(info.Mode())
}
