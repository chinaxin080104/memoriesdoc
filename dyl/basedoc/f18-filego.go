// 文件操作
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {
	//一次性读
	//main1801()
	//读取当前文件路径地址
	//main1802()
	//分片读
	//main1803()
	//带缓冲
	//main1805()
	//指定分割符
	//main1806()
	//文件复制
	//main1807()
	//目录操作
	main1808()
}

// 可以通过获取当前go文件的路径,然后用相对于当前go文件的路径去打开文件。
func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func main1801() {
	byteData, _ := os.ReadFile("study/hello.txt")
	fmt.Println(string(byteData))
}

func main1802() {
	file := GetCurrentFilePath()
	//输出 F:/dyl/basedoc/f18.go
	fmt.Println(file)
}

// 分片读
func main1803() {
	fmt.Println("start")
	file, _ := os.Open("study/hello.txt")
	defer file.Close()
	for {
		buf := make([]byte, 1)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		//按字母换行输出
		//fmt.Println(string(buf))
		//这里在for循环里面，调试不输出。
		fmt.Printf("%s", buf)
	}
	fmt.Println("end")
}

// 带缓冲读
func main1804() {
	file, _ := os.Open("study/hello.txt")
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}
}

// 指定分隔符
func main1805() {
	file, _ := os.Open("study/hello.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	//scanner.Split(bufio.ScanLines) //按照行读
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 文件写入
func main1806() {

	err := os.WriteFile("study/file1.php", []byte("<?php echo ='hello' ?>"), os.ModePerm)
	//返回 nil 代表没有错误
	fmt.Println(err)
	//常见的一些打开模式
	//如果文件不存在,就创建
	//os.O_CREATE | os.O_WRONLY
	//追加写
	//os.O_APPEND | os.O_WRONLY
	//可读可写
	//os.O_RDWR

	//const (
	//	O_RDONLY int = syscall.O_RDONLY // 只读
	//	O_WRONLY int = syscall.O_WRONLY // 只写
	//	O_RDWR   int = syscall.O_RDWR   // 读写
	//
	//	O_APPEND int = syscall.O_APPEND // 追加
	//	O_CREATE int = syscall.O_CREAT  // 如果不存在就创建
	//	O_EXCL   int = syscall.O_EXCL   // 文件必须不存在
	//	O_SYNC   int = syscall.O_SYNC   // 同步打开
	//	O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件
	//)

	//文件的权限 -主要用于linux 系统,代表文件的模式和权限位。
	//[x][y][z]
	//第一个是文件所有者所拥有的权限。
	//第二个是文件所在组对其拥有的权限。
	//第三个占位符是指其他人对文件拥有的权限。
	//每一位的权限，都是由读，写，执行相加得到。
	//R 读 Read缩写，八进制值为4
	//W 写 Write 八进制值为2
	//X 执行 Execute 八进制值为1
	//0444 表示三者均为只读的权限。
	//0666 表示三者均为读写权限。
	//0777 表示三者均为 读写执行的权限。
	//0764 表示所有者拥有 读写执行的权限（1+2+4） 组拥有 （2+4）的权限。 其他用户 （4）只读权限。
}

func main1807() {
	//io.Copy(dst Writer, src Reader) (written int64, err error)
	read, _ := os.Open("study/file1.php")
	write, _ := os.Create("study/file3.php") // 默认是 可读可写，不存在就创建，清空文件
	n, err := io.Copy(write, read)
	fmt.Println(n, err)
}

func main1808() {
	dir, _ := os.ReadDir("study")
	for _, entry := range dir {
		info, _ := entry.Info()
		//获取文件中文件的名字，大小。
		fmt.Println(entry.Name(), info.Size()) // 文件名，文件大小，单位比特
	}
}
