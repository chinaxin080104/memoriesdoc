// 反射
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	//判断一个变量是否是结构图，切片，map
	//main1901()
	//通过反射修改值
	//main1902()
	//结构体反射
	//main1903()
	//修改结构体中的某些值
	//main1904()
	//调用结构体中的方法。
	main1905()
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

type Student struct {
	Name string
	Age  int `json:"age"`
}

func main1903() {
	s := Student{
		Name: "小明同学",
		Age:  24,
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonField := field.Tag.Get("json")
		if jsonField == "-" {
			// 说明json的tag是空的
			jsonField = field.Name
		}
		fmt.Printf("Name: %s, type: %s, json: %s, value: %v\n", field.Name, field.Type, jsonField, v.Field(i))
		//Name: Name, type: string, json: , value: 小明同学  ---因为Name 定义时，json 为空。
		//Name: Age, type: int, json: age, value: 24
	}
}

type Teacher struct {
	Name1 string `big:"-"`
	Name2 string `json:"name_2"`
	Name3 string
}

func main1904() {
	s := Teacher{
		Name1: "张老师",
		Name2: "李老师",
		Name3: "董老师",
	}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(&s).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		bigField := field.Tag.Get("big")
		//判断类型是否是字符串
		if field.Type.Kind() != reflect.String {
			continue
		}
		if bigField == "-" {
			continue
		}
		//修改值。@修改的是 结构体中的类型。
		valueField := v.Field(i)
		valueField.SetString(strings.ToTitle(valueField.String()))
	}
	//{张老师 李老师 董老师}
	fmt.Println(s)
}

// 结构体方法。
func (Student) See(name string) {
	fmt.Println("see name:", name)
}

// 调用结构体方法。
func main1905() {
	s := Student{
		Name: "小林子",
		Age:  21,
	}
	//获取变量s 的类型信息
	t := reflect.TypeOf(s)
	//获取变量s 的值信息。
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumMethod(); i++ {
		methodType := t.Method(i)
		fmt.Println(methodType.Name, methodType.Type)
		//t.NumMethod() 获取结构体类型 Student 的方法数量，然后通过循环遍历这些方法。
		//t.Method(i) 获取第 i 个方法的 reflect.Method 类型信息。打印出方法的名字 methodType.Name 和方法的类型 methodType.Type。
		if methodType.Name != "See" {
			//在循环中，检查方法名是否为 "See"。如果不是 See 方法，则跳过继续检查下一个方法。
			continue
		}
		//调用See 方法
		methodValue := v.Method(i)
		//methodValue.Call() 用于调用方法。这里通过 Call 函数调用 See 方法，并传递参数。参数需要使用 reflect.Value 类型的切片进行传递，所以上面用了 reflect.ValueOf("雨林木风") 来将字符串转换为 reflect.Value
		methodValue.Call([]reflect.Value{
			reflect.ValueOf("雨林木风"), // 注意这里的 切片类型
		})
	}

}
