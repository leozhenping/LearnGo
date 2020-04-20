package main

/*

	条件变量本身不是锁， 经常与锁结合使用。

	结合条件变量使用锁的步骤:
		1. 判断条件变量
		2. 加锁
		3. 访问公共区
		4. 唤醒阻塞在条件变量上的对端.

	go标准库中的sync.Cond类型代表了条件变量，条件变量要与锁（互斥、读写锁）一起使用。成员变量L代表与条件变量搭配使用的锁。

	type Cond Struct {
		noCopy noCopy
		L Locker
		notify nofityList
		checker copyChecker
	}

	对应的有3个常用方法，wait, signal, Broadcast.

	使用流程:
		1. 创建条件变量:   var cond sync.Cond
		2. 指定条件变量使用的锁: cond.L = new(sync.mutex)
		3. 加锁: cond.L.Lock()， 给公共区加锁（互斥锁）, 互斥锁初始值为0, 表示未加锁状态。
		4. 判断是否到达阻塞条件。（缓冲区满. 或者缓冲区空）, -- 使用for循环判断
			for len(ch) == cap(ch) {
				cond.Wait() // 作用: 1: 阻塞，2：解锁，3： 加锁
			}
		5. 访问公共区 -- （读或写数据,或其他任务）
		6. 解锁条件变量用到的锁； cond.L.Unclock()
		7. 唤醒阻塞在条件变量上的对端。 。
 */


