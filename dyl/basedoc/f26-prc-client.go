// RPC 客户端
package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Num int
}

func main() {
	//rpc 客户端调用服务端操作。

	req := Req{100, 221}
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	var res Res
	//调用,然后取地址,函数功能在rpc 服务端实现。服务端返回实现结果。
	client.Call("Server.Add", req, &res)
	fmt.Println(res)
}
