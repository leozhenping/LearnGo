package main

/*
	死锁发生的情况:
		1. 单go程自己死锁:
			channel应该在至少2个以上的go程中进行通信。
		2. go程间channel访问顺序导致死锁
			使用channel一端读（写）， 要保证另一端写（读）操作，同时有机会执行。
		3. 多go程, 多channel交叉死锁
		4. 在go语言中， 尽量不要将互斥锁、读写锁与channel混用, 混用会产生隐性死锁。
*/

func main() {

	// 死锁情况一: 单go程自己死锁。
	/*

		ch := make(chan int)
		ch <- 789
		num := <- ch
		fmt.Println(num)

	*/

	// 死锁情况二: go程间channel访问顺序导致死锁
	/*

		ch1 := make(chan int)
		num1 := <- ch1
		fmt.Println(num1)
		go func() {
			ch1 <- 789
		}()

	*/

	// 死锁情况三:


	/*

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
				fmt.Println(num)
			}
		}
	}()
	for {
		select {
		case num := <- ch2:
			ch1 <- num
			fmt.Println(num)
		}
	}
	*/

}
