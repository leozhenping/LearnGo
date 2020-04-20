package main

import (
	"fmt"
	"runtime"
)

/*

	runtime.Gosched():  出让当前go程所占用的cpu时间片。
	runtime.Goexit():  结束调用该函数的当前goroutine.  在runtime.Goexit()调用之前注册的事件，将不受影响。
	runtime.GOMAXPROCS : 用来设置可以并行计算的CPU核数的最大值。并返回上一次调用成功的设置值。


*/

func main() {
	go func() {
		for {
			fmt.Println("son goroutine...")
		}
	}()
	for {
		fmt.Println("main goroutine...")
		runtime.Gosched() // 将cpu时间片出让出去后，go子程将获得更多的时间片。
	}
}
