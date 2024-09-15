// 单元测试 -对应测试文件, _test.go
// 测试模版需要引入testing库。
package main

import "fmt"

func Addfuc119(a int, b int) int {
	return a + b
}

func main() {
	sum := Addfuc119(1, 1111)
	fmt.Println(sum)
}
