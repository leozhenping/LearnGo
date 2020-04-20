package main

import (
	"fmt"
	"time"
)

/*
	timer: 创建定时器，指定定时时长，定时到达后， 系统会自动向定时器的成员C写系统当前时间(对chan的写操作).
	读取Timer.C得到定时后的系统时间，并且完成一次chan的读操作。

		type Timer struct {
			C <-chan Time
			r runtimeTimer
		}

*/

// 定时器3种定时方法
func main() {

	// 1. time.sleep
	time.Sleep(time.Second * 1) // 定时一秒

	// 2. time.timer
	fmt.Println("当前时间： ", time.Now())
	myTime := time.NewTimer(time.Second * 2)
	timeNow := <- myTime.C
	fmt.Println("现下时间: ", timeNow)

	// 3. time.after
	fmt.Println("当前时间： ", time.Now())
	myTime2 := time.After(time.Second * 2)
	timeNow = <- myTime2
	fmt.Println("现下时间: " , timeNow)


}
