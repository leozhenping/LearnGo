package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func someFunc() {
	ctx, cancel := context.WithCancel(context.Background())
	go doFunc(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

func doFunc(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Println("done")
			return
		default:
			logg.Println("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "Test ", log.Ltime)
	someFunc()
	logg.Println("exit")
}

/*
执行输出结果:
Test 12:17:37 work
Test 12:17:38 work
Test 12:17:39 work
Test 12:17:40 work
Test 12:17:41 work
Test 12:17:42 work
Test 12:17:43 work
Test 12:17:44 work
Test 12:17:45 work
Test 12:17:46 exit

为什么没有输出done, 但程序也退出了呢？原因:
当执行cancel()函数时，<-c.Done()已经被执行到了。 但此时主goroutine已经执行到了exit， 并退出了程序。所以done没有被显示出来。
*/
