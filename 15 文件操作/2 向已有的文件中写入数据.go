package main

import (
	"fmt"
	"os"
)

/*
打开文件:
	OpenFile()这个函数有三个参数:
	第一个参数表示: 打开文件的路径
	第二个参数表示: 模式，常见的模式有:
		O_RDONLY: 只读模式
		O_WRONLY: 只写模式
		O_RDWR: 可读可写模式
		O_APPEND: 追加模式

	第三个参数表示: 权限, 取值范围(0-7) // 遵循Linux权限的rwx规则.
	表示如下:
		0 : 没有任何权限
		1 : 执行权限（如果是可执行文件，是可以运行的）
		2 : 写权限
		3 : 写权限和执行权限
		4 : 读权限
		5 : 读权限与执行权限
		6 : 读权限与写权限
		7 : 读权限，写权限，执行权限

file, err := os.OpenFile("test",os.O_APPEND,6)

 */

func main() {
	var str string = "2222222222222222\n"
	file, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND,6)  // 单独只使用Append模式会报文件句柄错误。
	if err != nil {
		fmt.Println(err)
	}

	n, err1 := file.Write([]byte(str))
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(n)
	}

	defer file.Close()

}
