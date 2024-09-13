// 自定义函数
package main

import "fmt"

//为什么会有自定义类型
//自定义类型可以帮我们更好的抽象和封装数据.让代码更加 易读,易懂,易维护
//结构体 就是自定义类型中的一种.

type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = -1 //校验失败的错误
	ServiceErrCode Code = -2 //服务错误
)

func main() {
	//自定义类型
	//main0901()
	//类型别名
	main0902()
}

func (c Code) GetMsg() string {
	code := Code(c)
	var msg string
	switch code {
	case SuccessCode:
		msg = "success"
		break
	case ValidCode:
		msg = "check fail"
		break
	case ServiceErrCode:
		msg = "service error"
		break
	}
	return msg
}

func main0901() {
	//输出三种不同的状态。
	fmt.Println(SuccessCode.GetMsg())
	fmt.Println(ValidCode.GetMsg())
	fmt.Println(ServiceErrCode.GetMsg())
	//判断i,是否与successCode 相等。
	var i int
	fmt.Println(int(SuccessCode) == i)
}

type AliasCode = int
type MyCode int

const (
	SuccessCode2      MyCode    = 0
	SuccessAliasCode2 AliasCode = 0
)

// MycodeMethod 自定义类型可以绑定自定义方法。
//func (m MyCode) MyCodeMethod() {
//
//}

// MyAliasCodeMethod 类型别名，不可以绑定方法。
//func (m AliasCode) MyAliasCodeMethod() {
//
//}

func main0902() {
	// 类型别名，打印它的类型还是原始类型 main.MyCode int
	fmt.Printf("%T %T \n", SuccessCode2, SuccessAliasCode2)
	var i int
	fmt.Println(SuccessAliasCode2 == i)
	//SuccessCode2 注意这里的 SuccessCode2 为单独的类型。
	fmt.Println(int(SuccessCode2) == i)
}
