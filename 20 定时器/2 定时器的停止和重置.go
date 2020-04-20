package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 创建定时器
	mytimer := time.NewTimer(20 * time.Second)
	mytimer.Reset(1 * time.Second) // 重置定时器
	go func() {
		<- mytimer.C
		fmt.Println("现在时间: ", time.Now())
	}()

	mytimer.Stop() // 停止定时器

	for {
		;
	}


}