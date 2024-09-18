package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// handler
func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("恭喜您，你访问到了handerer @dyl"))
}
func main() {
	//浏览器访问 http://127.0.0.1:8080/index :恭喜您，你访问到了handerer @dyl
	//简单的访问
	//main2301()
	//
	main2302()
}

func main2301() {
	// 回调函数-调用handler
	http.HandleFunc("/index", handler)
	// 绑定服务
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		//res.Write([]byte("<h1>hello,world!</h1>"))
		//打印当前目录
		//程序运行目录=当前目录+引用目录。
		dir, _ := os.Getwd()
		fmt.Println("Current working directory:", dir)
		//获取不到文件。解决方法，找打服务区运行的根目录
		//<?php echo ='hello' ?>
		data, err := os.ReadFile("dyl/basedoc/study/file1.php")
		if err != nil {
			fmt.Println(data)
			fmt.Println(err)
		}
		res.Write(data)
	case "POST":
		// 获取body数据
		data, err := io.ReadAll(req.Body)
		// 拿请求头
		contentType := req.Header.Get("Content-Type")
		fmt.Println(contentType)
		//switch contentType {
		//case "application/json":
		//
		//}

		if err != nil {
			fmt.Println(data)
		}
		ma := make(map[string]string)
		json.Unmarshal(data, &ma)
		fmt.Println(ma["username"])

		type User struct {
			Username string `json:"username"`
		}
		var user User
		json.Unmarshal(data, &user)
		fmt.Println(user)
		// 设置响应头
		header := res.Header()
		header["token"] = []string{"y1gyf156sdgT%d44hjgj"}
		res.Write([]byte("hello 枫枫 POST"))
	}
}

func main2302() {
	// 1. 绑定回调 -调用IndexHandler
	http.HandleFunc("/index", IndexHandler)
	// 2. 注册服务
	fmt.Println("listen server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
