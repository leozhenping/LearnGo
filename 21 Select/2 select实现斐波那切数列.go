package main

import (
	"fmt"
	"runtime"
)

func fiboTest(fchannel <- chan int, ch <- chan bool)  {
	for {
		select {
		case num := <- fchannel:
			fmt.Print(num, " ")
		case <- ch:
			runtime.Goexit()
		}
	}

}


func main() {

	fibonacci := make(chan int)
	quit := make(chan bool)

	go fiboTest(fibonacci, quit)
	x,y := 1, 1
	for i:=0; i<20; i++ {
		fibonacci <- x
		x, y = y, x +y // 斐波那切数据的规律。
	}
	quit <- true

}

// 数列示例: 1 1 2 3 5 8 13 21 34 55 89 144 233