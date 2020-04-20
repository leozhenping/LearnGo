package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	go func() {
		for {
			readBuf := make([]byte, 4096)
			n, err := os.Stdin.Read(readBuf)
			if err != nil {
				fmt.Println(err)
				continue
			}
			conn.Write(readBuf[:n])
		}
	}()

	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)

		if n == 0 {
			fmt.Println("服务器端链接关闭。")
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("从服务器端读到数据: ", string(buf[:n]))
	}
}
