package main

/*

func ReadWrite() bool {
	file.Open("file")
	defer file.close() // 打开和关闭写在一起，方便管理。 还不容易遗漏。 defer总在函数执行最后执行。
	if aFailure {
		return  false
	} else if bFailure {
		return false
	}
}

*/

// 如果一个函数中有多个defer语句， 它们会以LIFO(后进先出)的顺序执行

