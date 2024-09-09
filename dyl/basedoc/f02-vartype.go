// 基础数据类型
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("基础数据类型")

	//整数型
	//funcuint()
	//funcfloat()
	funcchar()
	funcstring()
	moreline()
	burlfa()
	zerofaq()
}

func funcuint() {
	// var n1 uint8 = 2
	// var n2 uint16 = 2
	// var n3 uint32 = 2
	// var n4 uint64 = 2
	// var n5 uint = 2
	// var n6 int8 = 2
	// var n7 int16 = 2
	// var n8 int32 = 2
	// var n9 int64 = 2
	// var n10 int = 2
	//默认的数字定义是int类型
	//前面的u代表无符号,存正整数。
	//后面的数字是2进制的位数。
	//uint8 别名是 byte ,一个字节=8个比特位
	fmt.Printf("%.0f", math.Pow(2, 63))
	fmt.Println("")
	var n1 int = 9223372036854775807
	fmt.Println(n1)
	// var n2 int = 9223372036854775808 //数字超出限制
	// fmt.Println(n2)
}

func funcfloat() {
	const F1 = math.MaxFloat32
	const F2 = math.MaxFloat64
	fmt.Println(F1, F2)
}

// 字符
func funcchar() {
	var c1 = 'a'
	var c2 = 97
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Printf("%c %c\n", c1, c2) //以字符格式打印

	var r1 rune = '国' //字符的赋值为单引号。
	fmt.Printf("%c\n", r1)
}

// 字符串类型与转义字符
func funcstring() {
	var s string = "中国人"
	fmt.Println(s)

	fmt.Println("颜色\t红色")                  //制表符
	fmt.Println("回答\tyes")                 //回车
	fmt.Println("你好，\"中国人\"")              //双引号
	fmt.Println("这是测试,\r你知道吗")             //回到行首.这里光标会覆盖这是测试，然后输出你知道吗？
	fmt.Println("上传路径是:c:\\app\\main.exe") //反斜杠
}

// 多行字符串,在” 反引号 引号里面的字符会原样输出``````  在英文状态直接按esc 下面的键。就是了。
func moreline() {
	var s = `今天是\"hello"
	阳历的9月7号
	真好啊`

	fmt.Println(s)
}

// 布尔类型 只有true,和假
func burlfa() {
	// 布尔类型变量的默认值为false
	// Go 语言中不允许将整型强制转换为布尔型
	// 布尔型无法参与数值运算，也无法与其他类型进行转换
	var T = true
	var F = false
	fmt.Println(T)
	fmt.Println(F)
}

// 零值问题
func zerofaq() {
	// 如果我们给一个基本数据类型只声明不赋值
	// 那么这个变量的值就是对应类型的零值，例如int就是0，bool就是false，字符串就是""
	var a1 int
	var a2 float32
	var a3 float64
	var a4 bool
	var a5 string

	fmt.Printf("%#v\n", a1) //0
	fmt.Printf("%#v\n", a2) //0
	fmt.Printf("%#v\n", a3) //0
	fmt.Printf("%#v\n", a4) //false
	fmt.Printf("%#v\n", a5) //""
}
