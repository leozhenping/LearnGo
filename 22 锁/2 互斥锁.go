package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex // 创建一个互斥锁。新建的互斥锁状态为0， 0表示未加锁状态。

func printer(str string)  {
	mutex.Lock()
	for _, s := range str {
		fmt.Printf("%c",s)
		time.Sleep(time.Millisecond * 2)
	}
	mutex.Unlock()
}

func person1()  {
	printer("hello")

}

func person2()  {
	printer("world")

}

func main() {

	go person1()
	go person2()

	for {
		;
	}
}