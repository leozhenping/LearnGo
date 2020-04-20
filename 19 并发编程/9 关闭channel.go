package main

import (
	"fmt"
	"time"
)

/*
	close(channel) : 关闭channel
	数据不发送完，不关闭channel
	已经关闭的channel， 不能再向其写入数据。
	已经关闭的channel， 可以继续从中读取数据。读出来的数是类型的默认值。

	关闭无缓冲channel后，再读时： 返回默认值
	关闭有缓冲channel后，再读时，如缓冲区中有数据，先读数据。 数据读完后，返回默认值。
*/

func main() {

	channel := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		close(channel)
	}()

	time.Sleep(time.Second * 2)

	/*
		// 方法一 : 循环读入channel的数据
		for {
			if num, ok := <-channel; ok {
				fmt.Println("go主程读到: ", num)
			} else {
				fmt.Println("go程读入数据结束. 读取到的数据为: ", num) // 执行结果: num 为 0
				break
			}
		}
	*/

	// 方法二: 使用range读取channel的数据。
	for num := range channel {
		fmt.Println("go主程读到: ", num)
	}
}
