package main

import "fmt"

/*
	默认channel是双向的。

	单向channel分为:
		单向写channel :
			定义: var sendCh chan <- int , make(chan <- int)
		单向读channel
			定义: var recvCh <- chan int , make(<-chan int)

	转换:
		1. 双向channel可以隐式转换为任意一种意向channel
		2. 单向channel不能转换为双向channel。


	传参:
		传引用
 */

func writeCH(ch chan <- int, n int)  {
	ch <- n
}

func readCh(ch <- chan  int)  {
	num := <- ch
	fmt.Println("读到写入数据: ", num)
}

func main() {
	channel := make(chan int) // 双向ch

	go writeCH(channel, 88) // 由writeCh单向channel写入

	readCh(channel) // 由readCh单向读出


	/*
		// 此处为一个隐式转换，可以将双向channel转换为单向写、读channel
		var rch <- chan int = channel
		var wch chan <- int = channel
	 */

}
