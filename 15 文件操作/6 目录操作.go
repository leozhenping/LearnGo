package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 1. 打开目录
	/*
		fd, err := os.OpenFile("dir", os.O_RDWR, os.ModeDir)  // ModeDir是打开目录通常使用的权限

		fd是一个可权读写目录的文件指针


		目录操作常用函数：
			fd.Readdir(n int) ([]FileInfo, error)
				n : n为正整数时，表示读几个目录项，如果为-1表示读所有目录项
				FileInfo: 其内部保存了文件名，其是一个接口

					type FileInfo Interface {
						Name() string // 文件名
						Size() int64 // 文件长度
						Mode() FileMode // 文件权限
	                    ModTime() time.Time // 文件的修改时间
						IsDir() bool
						Sys() interface{}
					}
	 */
	// 查找目录下的所有dmg文件并删除。
	fd, err := os.OpenFile("/Users/weizhenping/Downloads", os.O_RDONLY, os.ModeDir)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
	}
	info, err := fd.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range info {
		if ! file.IsDir() {
			if strings.HasSuffix(file.Name(), "dmg") {
				os.Remove(file.Name())
			}
		}
	}
}
