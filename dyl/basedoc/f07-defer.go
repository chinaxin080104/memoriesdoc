// init 函数与defer 函数
package main

import "fmt"

//init（）函数特点
//不能被其他函数调用,而是在main函数执行之前，自动被调用。
//init函数不能作为参数传入.
//不能有传入参数与返回值.

//执行顺序如下：
//初始函数定义一
//初始函数定义二
//初始函数定义三
//初始函数定义四
//初始函数定义五
//main

func init() {
	fmt.Println("初始函数定义一")
}
func init() {
	fmt.Println("初始函数定义二")
}

func init() {
	fmt.Println("初始函数定义三")
}
func main() {
	//defer 函数 用于注册延迟调用
	//这些调用直到return 前才被执行。因此,可以用来做资源清理。
	//多个defer 语句,按先进后出 的方式执行。谁距return 近,谁先执行。 --先近后出。
	defer fmt.Println("defer4")
	deffunc()
	fmt.Println("defer3")
	fmt.Println("main")
}

func deffunc() {
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}

//func init() {
//	fmt.Println("初始函数定义四")
//}
//func init() {
//	fmt.Println("初始函数定义五")
//}
