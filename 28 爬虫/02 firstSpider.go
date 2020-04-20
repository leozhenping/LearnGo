package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func getUrl(url string , i int) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("err1: ", err1)
	}
	defer resp.Body.Close()
	buff := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buff)
		if err2 == io.EOF {
			fmt.Printf("%d读取完成。\n", i)
			break
		}
		if err2 != nil {
			err = err2
			fmt.Println("nil", err2)
			break
		}
		result += string(buff[:n])
	}
	return
}
func Working(start, end int, wg *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		go func(i int) {
			url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
			result, err := getUrl(url, i)
			if err != nil {
				fmt.Println(err)
			}
			fd, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
			if err != nil {
				fmt.Println(err)
			}
			fd.WriteString(result)
			fd.Close()

			wg.Done()
		}(i)
	}
}
func main() {
	var start, end int
	var wg sync.WaitGroup
	fmt.Println("请输入起始页面: ")
	fmt.Scan(&start)
	fmt.Println("请输入终止页面: ")
	fmt.Scan(&end)
	wg.Add(end)
	Working(start, end, &wg)
	wg.Wait()
}
