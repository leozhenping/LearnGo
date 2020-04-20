package main

import (
	"fmt"
	"io"
	"net"
	"regexp"
	"strings"
	"time"
)

type Client struct {
	Name   string
	Ipaddr string
	C      chan string
}

var message = make(chan string)

var clientMap = make(map[string]Client)

func handlerMessage() {
	for {
		msg := <-message
		for _, client := range clientMap {
			client.C <- msg
		}
	}

}

func WriteToClient(conn net.Conn, client Client) {
	for {
		msg := <-client.C
		conn.Write([]byte(msg))
	}

}

func MakeMsg(client Client, msg string) string {
	return "[" + client.Name + "]" + ": " + msg + "\n"
}

func HandlerConn(conn net.Conn) {
	defer conn.Close()
	remoteInfo := conn.RemoteAddr().String()
	client := Client{remoteInfo, remoteInfo, make(chan string)}
	clientMap[remoteInfo] = client

	message <- MakeMsg(client, "login")
	go WriteToClient(conn, client)

	buf := make([]byte, 1024)
	isQuit := make(chan bool)
	hasData := make(chan bool)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Println("conn.read", err)
				}
			}
			if n == 0 {
				isQuit <- true
			} else {
				msg := string(buf[:n-1])
				var clientList string
				if msg == "who" {
					conn.Write([]byte("Online Clinet list: \n"))
					for _, client := range clientMap {
						clientList += client.Name + "\n"
					}
					conn.Write([]byte(clientList))

				} else if status, _ := regexp.Match("rename.*", []byte(msg)); status {
					name := strings.Split(msg, "|")[1]
					client.Name = name
					clientMap[remoteInfo] = client
					conn.Write([]byte("用户名修改成功.\n"))

				} else {
					message <- MakeMsg(client, msg)
				}

			}
			hasData <- true
		}
	}()
	for {
		select {
		case <- isQuit:
			delete(clientMap, remoteInfo)
			message <- MakeMsg(client, "logout.")
			close(client.C) // 关闭WriteToClient的go程。
			return
		case <- hasData:
			// 什么也不做。只是每次在超时时间内，重置time.after定时器。
		case <- time.After(time.Second * 10):
			delete(clientMap, remoteInfo)
			message <- MakeMsg(client, "time out leave.")
			close(client.C) // 关闭WriteToClient的go程。
			return
		}
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	defer listener.Close()
	if err != nil {
		fmt.Println("net.listen", err)
	}

	go handlerMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept", err)
		}

		go HandlerConn(conn)
	}

}
