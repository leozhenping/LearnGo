package main

import "fmt"

// 模拟移动硬盘和U盘的读写数据

// 定义接口，包含读和写的功能
type Disk interface {
	Read()
	Write()
}

// 定义移动硬盘类
type Mdisk struct {

}

// 实现移动硬盘读取数据
func (m *Mdisk) Read()  {
	fmt.Println("移动硬盘读取数据。")
}

// 实现移动硬盘写入数据
func (m *Mdisk) Write()  {
	fmt.Println("移动硬盘写入数据。")
}

// 定义U盘类
type Udisk struct {

}

// 实现移动U盘读取数据
func (u *Udisk) Read()  {
	fmt.Println("u盘读取数据。")
}

// 实现移动U盘写入数据
func (u *Udisk) Write()  {
	fmt.Println("u盘写入数据。")
}

// 定义多态函数computer
func Computer(c Disk)  {
	c.Read()
	c.Write()
}

func main() {
	var u Udisk
	var m Mdisk

	Computer(&u)
	Computer(&m)
}
