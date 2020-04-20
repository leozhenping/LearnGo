package main

import (
	"fmt"
	"io"
	"os"
)

/*
读取文件步骤:
	1. 打开要读取的文件
	2. 对文件进行读取
	3. 关闭文件

	os.Open() : 其是调用OpenFile方法
	os.OpenFile()


	file.Read(b []byte)  // 返回读取的数据长度和错误信息
 */

func main() {
	/*
	// 当前案例不能循环读取文件，当文件的内容大小超过了buffer定义的大小，就不能正常显示了。 所以以下案例即解决此问题。
	fd, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
	}

	buffer := make([]byte, 1024*2) // 大小定义为2kb
	n,err := fd.Read(buffer) // 将读出来的数据，填充到buffer变量中。
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buffer[:n]))

	*/

	buffer := make([]byte, 10)
	fd, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(fd)
	}
	for {
		n, err := fd.Read(buffer)
		if err == io.EOF {  // io.EOF文件结尾。 当错误信息等于io.EOF就认为文件读取完毕了。
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}