// 关于打包部署
package main

//打包命令 go build
//go 项目的部署特别简单,编写完成后,只需要执行go build 即可打包为可执行文件。
//windows 下打包是exe文件。linux 下打包就是二进制文件。

//打当前目录下，xxx.go的包，这个包必须是一个main包。不然没有效果。
//go build xxxx.go  打当前目录下的mian包。注意只能有一个main 函数的包。

//强制对输出的文件进行重命名。
//go build -o main.exe  xxx.go

//交叉编译
//就是在windows 上，我开发的go程序，我也能打包为linux上的可执行程序。

//set 命令需要再cmd的命令行下执行.powershell 是无效的。
//set CGO_ENABLED=0
//set GOOS=linux
//set GOARCH=amd64
//go build -o main main.go

//CGO_ENABLED : CGO 表示 golang 中的工具，CGO_ENABLED=0 表示 CGO 禁用，交叉编译中不能使用
//CGO GOOS : 环境变量用于指定目标操作系统，mac 对应 darwin，linux 对应 linux，windows 对应 windows ，还有其它的 freebsd、android 等
//GOARCH：环境变量用于指定处理器的类型，386 也称 x86 对应 32位操作系统、amd64 也称 x64 对应 64 位操作系统，arm 这种架构一般用于嵌入式开发。比如 Android ， iOS ， Win mobile 等

//项目的根目录下写一个bat文件 然后放到linux服务器下，设置文件权限就可以直接运行了
// chmod +x main
//打包 web项目，配置文件和静态文件等 这些非go程序,是要一起复制到目标服务器里面的。
