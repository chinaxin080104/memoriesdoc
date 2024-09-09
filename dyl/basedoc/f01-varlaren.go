// 标准的变量定义,基础知识
package main

import "fmt"

var schoolName = "中华人民共和国。英文为 The People's Republic of China (PRC)"

func main() {
	// 先定义，在赋值
	var name string
	name = "心平气和"
	fmt.Println(name)
	//var 变量名 类型="变量值"
	var userEmail string = "735449187@qq.com"
	fmt.Println(userEmail)

	//省略变量类型
	var username = "dongyanglong"
	fmt.Println(username)

	//简短声明
	year := "2024"
	//如果一个变量定义了。但是没有赋值，那么这个变量的值就是这个类型的""零值"
	fmt.Println(year)

	//全局变量与局部变量。
	func2()
	//定义多个变量。
	func3()
	//常量定义
	func4()
	//格式化输出
	func5()
	//接受输入测试
	//func6()
}

//函数名 func2 的位置没有像C++ 里面那样,受到位置的制约。也就是说，只要存在。哪怕最后一行也行。
func func2() {
	var func2name = "函数名2"
	//在函数内定义的变量，必须要使用。
	fmt.Println(func2name)
}

//多个变量定义
func func3() {
	var name1, name2, name3 string //定义多个字符串类型的变量
	var a1, a2 = "1号出口", "2号出口"
	a3, a4 := "A栋", "B栋"
	//name1 declared and not used
	//a2 declared and not used
	//a4 declared and not used

	//变量批量定义与省略
	var (
		numberId int    = 01
		userName string = "小明同学"
		address         = "北京市"
	)

	//println可以像 var_dump()一样，连续输出多个对象。
	fmt.Println(name1, name2, name3, a1, a2, a3, a4)
	fmt.Println(numberId, userName, address)
}

//常量定义--定义的时候就要赋值，赋值之后就不要再修改了。
func func4() {
	const Chinesename string = "中国*"
	str := "我来自美丽的"
	str2 := "全称是:"
	//命名规范-核心思想首字母大写的变量，函数，方法，属性可以在包外进行访问。
	fmt.Println(str + Chinesename + str2 + schoolName)
}

//格式化输出

func func5() {
	fmt.Printf("%v\n", "你好啊")             //占位符输出
	fmt.Printf("%v %T\n", "英姿飒爽", "浩气长存") //输出字符的类型
	fmt.Printf("%d\n", 3)
	fmt.Printf("%.2f\n", 1.25)
	fmt.Printf("%s\n", "哈哈哈,下雨天")
	fmt.Printf("%#v\n", "") //这里是双引号

	//将格式化之后的内容，赋值给一个变量
	name := fmt.Sprintf("%v", "你好")
	fmt.Println(name)
}

//输入
func func6() {
	fmt.Println("请输入您的公司名称")
	var name string
	//这里其实就是c的指针，指到物理地址，寻指后,从地址中取到输入的值。
	fmt.Scan(&name)
	fmt.Println("您的公司是:", name)
}
