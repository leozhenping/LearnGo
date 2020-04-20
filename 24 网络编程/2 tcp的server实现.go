package main

import (
	"fmt"
	"net"
)

/*

	func Listen(network, address string) (listener, error)
		network：选用的协议, tcp, udp
		address: ip + port


	listener接口
		type Listener interface {
			Accept() (conn, error)
			Close() error
			addr() Addr
		}

	Conn接口:
		type Conn interface {
			Read(b []byte) (n int, err error)
			Write(b []byte) (n int, err error)
			Close() error
			LocalAddr() Addr
			RemoteAddr() Addr
			SetDeadline(t time.Time) error
			SetReadDeadline(t time.Time) error
			SetWriteDeadline(t time.Time) error
		}

 */

func main() {
	// 指定服务器通信协议，IP, Port
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// 阻塞监听客户端连接请求, 成功建立连接，返回用于通信的socket
	fmt.Println("服务器等待客户端建立连接...")
	conn, err := listener.Accept()
	defer conn.Close()
	if err !=nil {
		fmt.Println(err)
		return
	}

	fmt.Println("成功与客户端建立链接...")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(buf)

	fmt.Println(string(buf[:n]))

}
