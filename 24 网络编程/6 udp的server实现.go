package main

import (
	"fmt"
	"net"
	"time"
)

/*

	创建监听地址:
		func ResolveUDPAddr(network, address string) (*UDPAddr, error)
	创建用户通信的socket:
		func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
	接收Udp数据:
		func (c *UDPConn) ReadFromUDP(b []byte) (int, *UPDAddr, error)
	发送udp数据:
		func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)

 */

func main() {

	// 组织一个udp地址结构
	srvAddr, error := net.ResolveUDPAddr("udp","127.0.0.1:8003")
	if error != nil {
		fmt.Println("resolveUdpAddr error", error)
	}
	// 创建用于通信的socket
	udpConn, error := net.ListenUDP("udp", srvAddr)
	if error != nil {
		fmt.Println("ListenUdp error", error)
	}
	defer udpConn.Close()

	// 读取客户端发送的数据

	buf := make([]byte, 4096)
	// ReadFromUdp返回三个值: 读取到的字节数，客户端的地址， error
	n,  clientAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("readFromUdp error", err)
	}

	fmt.Printf("服务器读到%v的数据: %s", clientAddr, string(buf[:n]))

	// 回写数据给客户端
	dayTime := time.Now().String()
	_, error1 := udpConn.WriteToUDP([]byte(dayTime), clientAddr)
	if error1 != nil {
		fmt.Println("writeToUdp error", error1)
	}
}