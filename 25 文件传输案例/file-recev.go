package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, filename string)  {
	// 创建文件
	fd, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.create", err)
	}
	defer fd.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件接收完成.")
			return
		}
		if err != nil {
			fmt.Println("conn read", err)
			return
		}
		_, err = fd.Write(buf[:n])
		if err != nil {
			fmt.Println("fd write" ,err)
			return
		}
	}

}
func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net.listen ", err)
	}
	defer listener.Close()

	conn, accepterr := listener.Accept()
	if accepterr != nil {
		fmt.Println("listener accept ", accepterr)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("filename read ", err)
	}
	fileName := string(buf[:n])

	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("conn write", err)
	}

	recvFile(conn, fileName)

}
