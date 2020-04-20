package main

import (
	"fmt"
	"io"
	"os"
)

/*
创建文件:
	1. 导入os包，创建文件，读写文件的函数都是该包提供
	2。 指定创建的文件存放路径以及文件名
	3. 执行create()函数，进行文件创建
	4. 关闭文件


写入数据:
	1. WriteString  // 接受字符串类型，返回写入字符的长度和err, writeString底层是调用Write方法写入的。
		n, err := file.WriteString("string")
	2. Write        // 接受字节类型的切片, 返回写入的字节长度和err
		n, err := file.Write([]byte("string"))
	3. WriteAt      // 接受两个参数，向文件写入数据的内容，类型为[]byte, 第二个参数， 指定位置， 返回写入数据的长度和error
		n, err := file.WriteAt([]byte("string")),0) // 第二个参数表示指定位置

*/

func main() {

	// 创建文件
	file, err := os.Create("test1.go") // 返回文件指针和错误信息

	if err != nil {
		fmt.Println(err) // 文件创建错误，也需要关闭文件。以免超成不必要的内存开销。
	} else {
		// 写入数据

		// 1. WriteString写入数据
		n, err := file.WriteString("Hello world") // n: 表示数据的长度
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n) // 返回11, 表示写入的string的长度为11
		}
	}

	// 关闭文件
	defer file.Close()


	// 2.  Write写入数据
	var str string = "Hello World"
	file1, err1 := os.Create("test2.go")
	if err1 != nil {
		fmt.Println(err)
	} else {
		n1, err := file1.Write([]byte(str)) // []byte()将str转换为切片byte
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n1)
		}

	}
	defer file1.Close()



	// 3. WriteAt写入数据
	var str3 string = "hello world"
	var b3 string = "haha"

	file3, err3 := os.Create("test3")
	defer file3.Close()
	if err3 != nil {
		fmt.Println(err3)
	} else {
		num3, err:= file3.Write([]byte(str3))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("num3: ", num3)
		}
		/*
			io.SeekStart : 文件起始位置
			io.SeekCurrent: 文件当前位置
			io.SeekEnd: 文件结尾位置
		 */
		pointNumber,_ := file3.Seek(0, io.SeekEnd) // offset表示偏移多少个字符, io.SeekEnd 表示将光标移至文件最末尾并加上偏移的数量。返回文本长度和error
		number, err:=file3.WriteAt([]byte(b3), pointNumber)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(number)
	}
}
