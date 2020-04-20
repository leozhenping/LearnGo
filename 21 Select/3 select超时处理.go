package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(3 * time.Second): // 每轮for循环中， time.After都将重新计时。。
				quit <- true
				goto label  // 转跳至label标签位置继续执行。
			}
		}
	label: //标签功能只能在函数内部使用。
		print("finished....")
	}()

	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}

	<-quit
}
