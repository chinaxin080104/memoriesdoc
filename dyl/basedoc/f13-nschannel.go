// 异步模式下的channel,在协程函数里面写，在主线程里面接收数据
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//异步程序调用demo
	//main1301()
	main1302()
}

var moneyChan = make(chan int) //声明并初始化一个长度为0的信道。

// select 用
var moneyChan1 = make(chan int)    //声明长度为0的信道
var nameChan1 = make(chan string)  //同上
var doneChan = make(chan struct{}) //声明关闭信道
// 购物函数
func gopay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	moneyChan <- money
	wait.Done()
}

func main1301() {

	//小五 开始购物
	//小三 开始购物
	//小董 开始购物
	//小鱼 开始购物
	//小王 开始购物
	//小王 购物结束

	//5 true
	//小三 购物结束
	//1 true
	//小五 购物结束
	//2 true
	//小鱼 购物结束
	//4 true
	//小董 购物结束
	//3 true
	//0 false
	//moneyList []
	//购买完成 3.0348977s
	var wait sync.WaitGroup
	var moneyList []int
	startTime := time.Now()
	wait.Add(5)
	go gopay("小三", 1, &wait)
	go gopay("小五", 2, &wait)
	go gopay("小董", 3, &wait)
	go gopay("小鱼", 4, &wait)
	go gopay("小王", 5, &wait)

	go func() {
		defer close(moneyChan)
		//在协程函数里面等待上面三个协程函数结束。
		wait.Wait()
	}()

	//第一种写法
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		//跳出循环
	//		break
	//	}
	//}

	//第二种写法。如果for循环被close 关闭，for循环会自己结束。
	//moneyList [5 1 2 4 3]
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}
	fmt.Println("moneyList", moneyList)

	time.Sleep(2 * time.Second)
	fmt.Println("购买完成", time.Since(startTime))
}

// 声明购物方法
func nssend(name string, mone int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan1 <- mone
	nameChan1 <- name
	wait.Done()
}

// select 方法使用
func main1302() {
	//如果一个协程函数,往多个channel里面写东西，主线程数据获取？
	//go 使用select 异步从多个channel里面获取数据。
	var wait sync.WaitGroup
	startTime := time.Now()
	wait.Add(3)
	go nssend("张三", 21, &wait)
	go nssend("王五", 22, &wait)
	go nssend("李四", 23, &wait)

	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		wait.Wait()
	}()
	var moneyList []int
	var nameList []string

	//匿名函数
	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:
				return
			}
		}
	}
	event()
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)
}
