package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	fmt.Println(resp.Header)
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.ContentLength)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close() // 此处是关闭resp.body，而不是关闭resp。 resp.Close()是判断是否关闭，返回bool类型

	buf := make([]byte, 1024)
	var htmlString string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成。")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.body.read error ", err)
			break
		}

		htmlString += string(buf[:n])
	}
	fmt.Println(htmlString)
}
