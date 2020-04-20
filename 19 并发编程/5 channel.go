package main

import "fmt"

/*
	channel是go语言中的核心型，可以把它看成管道。 并发核心单元通过它就可以发送或者接收数据进行通讯，
	主要用来解决协程的同步问题。以及go程之间数据共享（数据传递）问题。
	goroutine运行在相同的地址空间，因此访问共享内存必须做好同步，
	goroutine奉行通过通信来共享内存，而不是共享内存来通信。
	引用类型channel可用于多个goroutine通讯。 其内部实现了同步，确保并发安全。

	channel的零值也是nil

	创建
		make(chan TYPE, capacity) // TYPE: 在channel中传递的数据类型, capacity: 容量， 0表示无缓冲channel， 大于0的值，表示有缓冲channel. 存满将阻塞。

	channel有两端:
		一端: 写入端, ,传入端. 语法表示: channel <-
		一端；读端, 传出端, 语法表示: <- channel
		要求： 读端和写端必须同时满足条件，才能在chan上进行数据流动。否则阻塞

 */

func main() {
	// 1. 定义channel
	channel := make(chan int)
	fmt.Println(channel)

	/*
		len(channel) : channel中剩余未读取数据个数。
		cap(channel) : 通道容量
	*/

}