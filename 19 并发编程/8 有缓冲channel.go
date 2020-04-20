package main

import (
	"fmt"
	"time"
)

/*
	有缓冲: 异步通信
	有缓冲channel， 容量为非0，
	len(channel) : channel中剩余未读数据个数。
	cap(channel) : 通道容量

	缓冲区可以进行数据存储。存储至容量上限才阻塞。

	有缓冲channel具备异步的能力, 不需要同时操作channel缓冲区。
 */

func main() {
	channel := make(chan int, 3)

	go func() {
		for i:=0; i<10; i++ {
			channel <- i
			fmt.Println("子go程: ", i)
		}
	}()

	time.Sleep(time.Second * 2)
	for x:=0; x<10; x++ {
		num := <- channel
		fmt.Println("主go程: ", num)
	}
}
