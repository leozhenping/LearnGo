package main

import (
	"context"
	"log"
	"os"
	"time"
)
var logg *log.Logger

func someFunc()  {
	/*
	withTimeout等价于withDeadline。
	 */

	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	go doFunc(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

func doFunc(ctx context.Context)  {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <- ctx.Done():
			logg.Println("done.")
			return
		default:
			logg.Println("work.")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "Test ", log.Ltime)
	someFunc()
	logg.Println("exit.")
}

/*
执行输出结果:
Test 12:24:55 work.
Test 12:24:56 work.
Test 12:24:57 work.
Test 12:24:58 work.
Test 12:24:59 done.

由于withTImeout只有5秒，当到达5秒后，c.done()就会被执行到了。
*/
