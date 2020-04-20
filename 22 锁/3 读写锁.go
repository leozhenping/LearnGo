package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex
var num int

func readGo(n int)  {
	for {
		rwMutex.RLock()
		fmt.Printf("第%d号go程，读取数据: %d\n", n, num)
		rwMutex.RUnlock()
	}
}


func writeGo (n int) {
	for {
		ranId := rand.Intn(999)
		rwMutex.Lock()
		num = ranId
		fmt.Printf("第%d号go程，写入数据: %d\n", n, num)
		rwMutex.Unlock()
	}
}

func main() {

	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())


	for i:=0;i<5 ;i++  {
		go readGo(i)
	}


	for x:=0; x<5; x++ {
		go writeGo(x)
	}

	for {
		;
	}
}
