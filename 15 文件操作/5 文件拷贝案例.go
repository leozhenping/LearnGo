package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	buffer := make([]byte, 1024)
	SrcFile, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
	}

	DstFile, err := os.Create("d.txt")
	if err != nil {
		fmt.Println(err)
	}

	for {
		n, err := SrcFile.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		DstFile.Write(buffer[:n])
	}

	defer SrcFile.Close()
	defer DstFile.Close()

}
