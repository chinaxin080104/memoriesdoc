// 反射
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//判断一个变量是否是结构图，切片，map
	//main1901()
	//通过反射修改值
	main1902()
}

func refType(obj any) {
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	//去判断具体的类型
	switch typeObj.Kind() {
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Struct:
		fmt.Println("struct")
	case reflect.String:
		fmt.Println("字符串")
	}
}

func main1901() {
	//结构体判断
	refType(struct {
		Name string
	}{Name: "小风范"})

	name := "小风"
	//字符串判断
	refType(name)
	//切片判断
	//refType([]string{"小风"})
}

// 通过反射修改值。
func refSetValue(obj any) {
	value := reflect.ValueOf(obj)
	elem := value.Elem()
	switch elem.Kind() {
	case reflect.String:
		fmt.Println("测试2")
	}
}

func main1902() {
	name := "峰峰"
	refSetValue(&name)
	fmt.Println(name)
}
