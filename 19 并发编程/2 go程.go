package main

import "fmt"

/*
	go语言中的并发程序主要使用两种手段来实现，goroutine和channel

	什么是goroutine?
			goroutine是Go语言并行设计的核心，有人称之为go程。 goroutine说到底其实就是协程。它比线程更小，十几个goroutine可能体现在底层就是五六个线程。
			go语言内部实现了这些goroutine之间的内存共享。 执行goroutine只需要极少的栈内存（大概4-5kb), 当然会根据相当数据伸缩。 也正因为如此，可同时运行成千上万个并发任务。
			goroutine比thread更易用，更高效，更轻便。

	go程创建于进程中。
	直接使用go关键字，放置于函数调用前面，产生一个go程。


	go程的特性:
		主go程结束，子go程随之退出。


 */

func sing()  {
	fmt.Println("唱歌...")

}

func main() {

	// 1. go程的创建
		go sing()

}
