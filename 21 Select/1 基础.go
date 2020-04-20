package main

import (
	"fmt"
)

/*
	go里面提供了一个关键字select, 通过select可以监听channel上的数据流动。
	select的用法与switch语法非常类似, 由select开始一个新的选择块，每个选择条件由case语句来描述。
	与switch语句相比，select有比较多的限制，其中最大的一条限制就是每个case语句里面必须是一个IO操作。

	每一个select语句中， go语言会按顺序从头至尾评估每一个发送和接收的语句
	如果其中的做生意一语句可以继续执行（即没有被阻塞）， 那么就从那些可以执行的语句中任意选择一条来使用。
	如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况:
		1. 如果给出了default语句， 那就会执行default语句， 同时程序的执行会从select语句后的语句中恢复。
		2. 如果没有default语句，那么select语句将阻塞，直到至少有一个通信可以进行下去。

	如果使用default语句，select就是一个忙轮循, 比较浪费CPU资源 。
	如果不人一用delfault语句，那么select就会一直阻塞，到时有信号进入。比较节省cpu资源。
*/

func main() {

	quit := make(chan bool)
	channel := make(chan int)
	channel2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		quit <- true
		close(channel)

	}()

	for {
		select {
		case num := <-channel:
			fmt.Println("读到channel1数据: ", num)
		case num := <-channel2:
			fmt.Println("读到channel2数据: ", num)
		case <-quit:
			return
			//runtime.Goexit() // 此处不能使用Goexit()， Goexit()只用于子go程中。 不能使用主go程中。
			//break // 只能跳出当前select中的case，不能跳出for循环。
		}
	}

}

/*
	注意事项:
		1. 监听的case中，没有满足监听条件将阻塞
		2. 监听的case中，有多个满足监听条件，任选一个执行。
		3. 可以使用default来处理所有case都不满足监听条件的状况。通常不用（会产生忙轮询)
		4. select自身不带有循环机制， 需借助外层for来循环监听。
		5. break只能跳出select , 类似于switch语句中的用法.


*/
