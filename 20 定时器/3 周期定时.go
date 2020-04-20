package main

import (
	"fmt"
	"time"
)

/*

	type Ticker struct {
		c <- chan Time
		r runtimeTimer
	}

	1. 创建周期定时器 myticker := time.NewTicker(time.second)
		定时时长到达后，系统会自动向Ticker的C中导入系统当前时间， 并且，每隔一个定时时长后，循环写入系统当前时间
	2. 在子go程中循环读取C, 获取系统写入的时间



 */
func main() {
	quit := make(chan bool)
	num := 0


	fmt.Println("系统时间为:  ", time.Now())
	timeTicker := time.NewTicker(time.Second)

	go func() {
		for {
			nowTime := <- timeTicker.C
			fmt.Println("当前时间为: ", nowTime)
			num ++

			if num == 8 {
				quit <- true
			}
		}
	}()

	<- quit
}