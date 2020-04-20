package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)


func HandlerConnect(c net.Conn) {
	defer c.Close()
	// 获取连接的客户端 addr
	addr := c.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")

	// 循环处理用户请求
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)

		// 客户端异常关闭，服务器处理操作
		if n == 0 {
			fmt.Println("服务器检测到客户端关闭，断开连接")
			return
		}

		// exit命令退出
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "EXIT\n" {
			fmt.Println("客户端正常退出")
			return
		}

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("服务器读到数据: ", string(buf[:n]))
		// 小写转大写， 回发给客户端

		//c.Write([]byte(strings.ToUpper(string(buf[:n]))))
		c.Write(bytes.ToUpper(buf[:n]))
	}

}

func main() {

	// 1. 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listener.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 2. 监听客户端连接请求
	for {
		fmt.Println("服务端接受客户端请求中...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		// 具体完成服务器与客户端的数据通信

		go HandlerConnect(conn)
	}

}
