// 网络编程
package main

import (
	"fmt"
	"io"
	"net"
)

//网络编程 包括
//TCP
//HTTP
//RPC
//webSocket

// 服务端
// 客户端
func main() {
	//运行后结果
	//	127.0.0.1:56966 进来了
	//	GET / HTTP/1.1
	//  Host: 127.0.0.1:8080
	//  Connection: keep-alive
	//	sec-ch-ua: "Not)A;Brand";v="99", "Opera";v="113", "Chromium";v="127"
	//	sec-ch-ua-mobile: ?0
	//	sec-ch-ua-platform: "Windows"
	//	Upgrade-Insecure-Requests: 1
	//	User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0
	//  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
	//	Sec-Fetch-Site: none
	//	Sec-Fetch-Mode: navigate
	//	Sec-Fetch-User: ?1
	//	Sec-Fetch-Dest: document
	//	Accept-Encoding: gzip, deflate, br, zstd
	//	Accept-Language: zh

	// 创建tcp的监听地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	// tcp监听
	listen, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		// 等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		// 获取客户端的地址
		fmt.Println(conn.RemoteAddr().String() + " 进来了")
		// 读取客户端传来的数据
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			// 客户端退出
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr().String() + " 出去了")
				break
			}
			fmt.Println(string(buf[0:n]))
		}

	}

}
