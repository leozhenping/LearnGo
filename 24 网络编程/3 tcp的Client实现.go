package main

import (
	"fmt"
	"net"
)

func main() {

	// 指定服务器通信协议， IP+ port创建通信套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 客户端主动写数据给服务端
	conn.Write([]byte("Are you Ready.!"))

	// 接收服务器发回的信息
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf[:n]))

}
