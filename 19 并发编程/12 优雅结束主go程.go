package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go func() {
		fmt.Println("go程1输出...")
		wg.Done() // // 操作完成，减少一个计数
	}()

	go func() {
		fmt.Println("go程2输出")
		wg.Done()
	}()

	wg.Wait() // 等待，直到计数为0
}
