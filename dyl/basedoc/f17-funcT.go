// 泛型函数
package main

import (
	"encoding/json"
	"fmt"
)

// 1.18版本后,go支持泛型。
func main() {
	//泛型函数
	main1701()
	//泛型结构体
	main1702()
	//泛型切片
	main1703()
	//泛型的map
	main1704()
}

func add[T int | float64 | int32](a, b T) T {
	return a + b
}

func main1701() {
	//因为有泛型的加入，m的类型。可以根据函数的泛型返回值在来定义。
	m := add(3.12, 4.12)
	fmt.Println(m)
}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main1702() {
	//调用结构体
	type User struct {
		Name string `json:"name"`
	}

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user := Response[User]{
		Code: 0,
		Msg:  "成功",
		Data: User{
			Name: "李小龙",
		},
	}

	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
	//不能这样用？
	userInfo := Response[UserInfo]{
		Code: 0,
		Msg:  "成功",
		Data: UserInfo{
			Name: "张三丰",
			Age:  24,
		},
	}
	byteData, _ = json.Marshal(userInfo)
	fmt.Println(string(byteData))

	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"小吴"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)
	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"小张","age":24}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)
}

// 泛型的切片
type MySlice[T any] []T

func main1703() {
	var mySlice MySlice[string]
	mySlice = append(mySlice, "西安")
	var intSlice MySlice[int]
	intSlice = append(intSlice, 2)
	fmt.Println(intSlice, mySlice)
}

// 泛型map
type MyMap[K string | int, V any] map[K]V

func main1704() {
	var myMap = make(MyMap[string, string])
	myMap["author"] = "李白"
	myMap["title"] = "登鹳雀楼"
	myMap["time"] = "唐"
	myMap["content"] = "白日依山尽，黄河入海流，欲穷千里目，更上一层楼"
	fmt.Println(myMap)
}
