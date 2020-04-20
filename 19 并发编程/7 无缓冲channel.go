package main

import "fmt"

/*
	无缓冲： 同步通信，要求两端同时在线
	无缓冲的channel应用于两个go程中，一个读，一个写。
	无缓冲的channel具备同步的能力。需要同步读写.
 */

func main() {

	var channel = make(chan int)

	go func() {
		for i := 1; i < 5; i++ {
			fmt.Println("子go程: " , i)
			channel <- i
		}
	}()
	for x := 1; x < 5; x++ {
		num := <-channel
		fmt.Println("主go程: ", num)
	}
}

/*
输出结果:
子go程:  1
子go程:  2  // 子go程连续打印两次，是因为主go程需要与子go程抢标准输出。此处子程的速度比主go程要快。
主go程:  1
主go程:  2
子go程:  3
子go程:  4
主go程:  3
主go程:  4
 */