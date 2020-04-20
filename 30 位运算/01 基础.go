package main

import "fmt"

/*
	位运算没有进位的概念。

    在位运算时， 将1参考为true, 0参考为false. 那么位与和位或的结果就与逻辑与&&和逻辑或||保持一致。

    & : 都为真是，结果为真
    &   1   0
    1   1   0
    0   0   1

    | : 一个为真，结果为真。
    |   1   0
    1   1   1
    0   1   0

    ^ : 相同为假，不同为真。
    ^   1   0
    1   0   1
    0   1   0

	位移运算示例:

	假定1为int8类型，1的二进制如下:
    0000 0001
       << 3
    000 0 0001 000 (最高丢弃3位，低位补3个0)，
		计算结果为: 0 0001 000, 十进制为: 8

    演算:  19 >> 2
        0001 0011
        00 0001 00
        计算结果为: 0000 0100 , 十进制为: 4
*/

const (
	LOG_DEBUG = 1 << iota // 此时二进制表示: 0000 0001 , 对应十进制为: 1
	LOG_INFO              //向左移位, 此时二进制表示: 0000 0010 , 对应十进制为: 2
	LOG_WARN              //向左移位，此时二进制表示: 0000 0100, 对应十进制为: 4
	LOG_ERROR             //向左移位，此时二进制表示: 0000 1000, 对应十进制为: 8
)

func main() {

	fmt.Println(LOG_DEBUG, LOG_INFO, LOG_WARN, LOG_ERROR) // 输出结果: 1 2 4 8
	LogLevel := 0
	// 1. 定义日志级别。
	// 需要开启警告和信息级别的日志
	LogLevel = LogLevel | LOG_WARN | LOG_INFO
	/*
		运算过程:
			LogLevel : 0000 0000
			LOG_WARN : 0000 0100
			LOG_INFO : 0000 0010

			0000 0000
			---------	或运算
			0000 0100
			---------
			0000 0100 与LOG_WARN位或结果.
			--------	或运算
			0000 0010
			---------
			0000 0110 : 与LOG_INFO位或结果: 对应十进制为: 6
	*/


	fmt.Println(LogLevel) // 输出结果为: 6

	// 再次开启信息级别的日志
	LogLevel |= LOG_INFO  // 0110 | 0010 , 或的结果为: 0110.
	fmt.Println(LogLevel) // 输出结果为: 6

	// 2. 检测日志级别
	// 检测某位是1还是0. 当拥有某个loglevel后，想确定哪个日志级别开启，如何处理？，使用位与运算即可。
	// 检测日志级别是否开启，示例如下: (返回true即为开启了对应级别的日志)
	// 位与的结果大于0表示对应位原本为1, 等于0表示对应位原本为0, 通过与0的比较，得到布布值。

	fmt.Println("debug", LogLevel & LOG_DEBUG, LogLevel & LOG_DEBUG != 0) // 计算过程: 0000 0110 & 0000 0001 , 结果为: 0000 0110
	fmt.Println("info", LogLevel & LOG_INFO, LogLevel & LOG_INFO != 0) // 计算过程: 0000 0110 & 0000 0010, 结果为: 0000 0010
	fmt.Println("warn", LogLevel & LOG_WARN, LogLevel & LOG_WARN != 0) // 计算过程: 0000 0110 & 0000 0100, 结果为: 0000 0100
	fmt.Println("error", LogLevel & LOG_ERROR, LogLevel & LOG_ERROR != 0) // 计算过程: 0000 01100 & 0000 1000, 结果为: 0000 0000


	// 3. 关闭日志级别
	// 核心思想: 使用反转运算。
	// 关闭info级别日志
	LogLevel = LogLevel ^ LOG_INFO // 计算过程: 0000 0110 ^ 0000 0010 , 结果为: 0000 0100
	fmt.Println("LogLevel", LogLevel) // 计算结果: 4

}

