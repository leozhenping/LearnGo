package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 50; i++ {
		fmt.Println("美女在唱歌....")
		time.Sleep(100 * time.Millisecond)
	}
}

func dance() {
	for i := 1; i < 50; i++ {
		fmt.Println("美女在跳舞")
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go sing()
	go dance()
	for {

	}

}
