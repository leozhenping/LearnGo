package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond // 定义Cond条件变量

func product(out chan<- int, idex int) {
	for {
		cond.L.Lock() // 获取锁
		for len(out) == 5 {
			cond.Wait() // 挂起并等待，阻塞并解锁(这两步是一个原子操作), 如果被对端唤醒，会再次加锁。
		}
		num := rand.Intn(999)
		out <- num
		fmt.Printf("%d号生产者生产了: %d \n", idex, num)
		cond.L.Unlock() // 解锁
		cond.Signal()   // 唤醒监听消息的对端
		time.Sleep(time.Millisecond * 200)
	}

}

func consumer(in <-chan int, idex int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("%d号消费者消费了: %d\n", idex, num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {

	cond.L = new(sync.Mutex) // 定义锁，使用互斥锁, 也可叫互斥量
	channel := make(chan int, 5)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go product(channel, i)
	}

	for i := 0; i < 5; i++ {
		go consumer(channel, i)
	}

	for {
		;
	}
}
