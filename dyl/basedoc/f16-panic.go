// 异常处理
package main

import (
	"fmt"
	"runtime/debug"
)

func init() {

	//获取配置文件路径报错,panic 会中断执行。
	//_, err := os.ReadFile("./doc/F:/doc/F.md")
	//if err != nil {
	//	panic(err.Error())
	//	//errorstr := err.Error()
	//	//fmt.Println(errorstr)
	//}
}

// 获取错误
func read() {

	//这个用于捕获异常的defer的延迟函数可以在调用链路上的任何一个函数上。
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //捕获异常,打印错误信息
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) //肯定会有一个panic 输出、
}

func main() {
	read()
	fmt.Println("hello,world")
	main1601()
}

// 遇到错误直接停止程序
func main1601() {
	fmt.Println("hello,dyl")
}
