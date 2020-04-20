package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, fullpath string)  {
	fd, err := os.Open(fullpath)
	if err != nil {
		fmt.Println("os.open", err)
	}
	defer fd.Close()

	buf := make([]byte, 4096)
	for {
		n, err := fd.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完成.")
			} else {
				fmt.Println("fd.read" ,err)
			}
			return
		}

		//conn.Write(buf[:n])
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("fd.write", err)
		}
	}

}

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Println("用户参数错误，请重新传参。")
		return
	}
	fullPath := args[1]
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		fmt.Println("os.stat", err)
	}
	fileName := fileInfo.Name()


	// 向服务端建立链接

	conn, err := net.Dial("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net.dial", err)
	}
	defer conn.Close()

	conn.Write([]byte(fileName))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.read", err)
	}

	if string(buf[:n]) == "ok" {
		sendFile(conn, fullPath)
	}

}
