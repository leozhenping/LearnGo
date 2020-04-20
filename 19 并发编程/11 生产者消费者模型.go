package main

import "fmt"

/*
	生产者:
	消费者:
	缓冲区:
		优点:
		1. 解耦： 降低生产者和消费者的偶合度
		2. 提高并发能力: 可以多个消费者、生产者同时写入和读出数据。
		3. 数据缓存： 生产者和消费者数据处理速度不一致时，可暂存数据。
 */

// 订单结构体
type Order struct {
	id int
}

// 生产订单
func producer(ch chan <- Order)  {
	for i:=1; i<10; i++ {
		order := Order{id:i + i}
		ch <- order
	}
	close(ch)
}

// 消费订单
func consumer(ch <- chan Order)  {
	for order :=  range ch {
		fmt.Println("接到订单, 单号为:  ",order.id)
	}
	
}

func main() {
	channel := make(chan Order, 5)
	go producer(channel)
	consumer(channel)
}
