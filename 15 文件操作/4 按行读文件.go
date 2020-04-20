package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fd, err := os.OpenFile("test.txt", os.O_RDWR, 7)
	defer  fd.Close()
	if  err != nil {
		fmt.Println(err)
	}

	// 创建一个带有缓冲区的reader
	reader := bufio.NewReader(fd)
	for {
		b, err := reader.ReadBytes('\n') // 以\n为结尾符，读取一行数据, 其返回bytes格式的数据和错误信息
		if err != nil {
			fmt.Println(err)
			return
		} else  if err == io.EOF {
			return
		}
		fmt.Print(string(b))
	}
}
