package main

import (
	"errors"
	"fmt"
)

// panic : 引发异常， 从而强制终止整个程序的执行。


// error接口

/*
error包的使用:
	1. 导入errors包
	2. 通过errorrs.New()函数描述错误信息
	3. 将错误返回
	4  接收返回的错误信息，并且进行判断
 */

func Test(num1 int, num2 int) (result int, err error)  {  // err默认为nil, error是一个接口类型。
	if num2 == 0 {
		err = errors.New("除数不能为0")
		return
	}
	result = num1 / num2
	return

}

func main() {
	number, err := Test(20, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number )
	}
}