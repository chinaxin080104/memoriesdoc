// 函数与指针引用
package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//函数的定义
	// sayWelcome()
	//函数参数
	// addfunc(1, 2)
	//函数参数2
	// morefunc(1, 2.2, 3.4, 3.1415926)

	//函数的返回值
	//fun1 := fun1() //fun1() (no value) used as value
	//fmt.Println(fun1)

	// fun2 := fun2() //输出1
	// fmt.Println(fun2)

	//多个返回值
	// num1, string1 := fun3()
	// fmt.Println(num1, string1) //1 错误

	// res := fun4()
	// fmt.Println(res) //abcd 字符串返回

	//匿名函数
	// nimingfunc()

	//高阶函数-根据用户输出的不同，执行不同的操作。
	// login()

	//高阶函数-抽离后。
	// loginv2()

	//函数的闭包规则
	// s := awaitAdd(2)(1, 2, 3)
	// fmt.Println(s)

	//值传递与引用传递
	funcyinyong()
}

func sayWelcome() {
	fmt.Println("Welcome!")
}

// 带参数函数定义.
func addfunc(a1 int, a2 int) {
	//fmt.Println(a1, a2)
	sum := a1 + a2
	//返回最终结果
	fmt.Println(sum)
}

// 多参数
func morefunc(numlist ...float32) {
	fmt.Println(numlist)
}

// 无返回值
func fun1() {
	return //无返回值
}

// 单返回值
func fun2() int {
	return 1
}

// 多返回值
func fun3() (int, error) {
	return 1, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
	//return
	return "abcd"
}

// 匿名函数
func nimingfunc() {
	var sum = func(x, y int) int {
		return x * y
	}
	//变量也可以是函数。傻傻的。
	fmt.Println(sum(9, 9))
}

// 多条件执行多任务 v1
func login() {
	// 请输入您需要执行的操作:

	//     1:登录
	//     2:个人中心
	//     3:注销
	fmt.Println("请输入您需要执行的操作:")
	fmt.Println(`
	1:登录
	2:个人中心
	3:注销`)
	var num int
	fmt.Scan(&num)
	var funcMap = map[int]func(){
		1: func() {
			fmt.Println("welcome 登录")
		},
		2: func() {
			fmt.Println("个人中心")
		},
		3: func() {
			fmt.Println("注销")
		},
	}
	funcMap[num]()
}

// 多条件执行多任务提取功能 v2
func login1() {
	fmt.Println("登录")
}
func userCenter() {
	fmt.Println("个人中心")
}
func logout() {
	fmt.Println("注销")
}

func loginv2() {
	// 请输入您需要执行的操作

	// 1.登录
	// 2.个人中心
	// 3.注销
	fmt.Println("请输入您需要执行的操作")
	fmt.Println(`
	1.登录
	2.个人中心
	3.注销
	`)
	var num int
	fmt.Scan(&num)
	var funcMap = map[int]func(){
		1: login1,
		2: userCenter,
		3: logout,
	}
	funcMap[num]()
}

// 函数的闭包
func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	//这里return 返回函数是因为返回类型里面约束了必须是函数。
	return func(i ...int) int {
		var sum int
		// _ 表示不接收i 返回的key
		for _, j := range i {
			sum += j
		}
		return sum
	}
}

func yingyong1(num int) {
	fmt.Print(&num) //输出year 的物理地址--// 0xc00000a0882024
	num = 2000
	//值传递，在函数体内，不管如何修改值，都不会影响外面的值。

}
func yingyong2(num *int) {
	//引用传递，我们需要直接将变量的内存地址传递进去。eg:2
	// &是取地址，*是解引用，去这个地址指向的值
	fmt.Println(num) //内存值
	*num = 2025      //这里的修改会影响外面的num
}

func funcyinyong() {
	year := 2024

	fmt.Println(&year) // 0xc00000a080
	yingyong1(year)
	fmt.Println(year)

	//引用传递
	yingyong2(&year)
	fmt.Println(year)
}
