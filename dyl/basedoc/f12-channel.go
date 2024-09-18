// 协程和channel
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//协程任务执行
	main1201()
	//主线程等待协成任务
	main1202()
	//协程间通信。 协程1传递数据给协成2,或者给主进程。
	//channel
	main1203()
	//go 通道的取值
	main1204()
}

func singer() {
	fmt.Println("小林,唱了一首歌-七律")
	fmt.Println("小黄,唱了一首歌-go going")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌比赛结束,小林是第一")
}

// 在go 中开启一个携程是非常简单的.
func main1201() {
	//Goroutine是Go运行时管理的轻量级线程
	go singer()
	go singer()
	go singer()
	go singer()

	time.Sleep(2 * time.Second)
	//如果我把这个主线程中的延时去掉之后，你会发现程序没有任何输出就结束了
	//这是为什么呢
	//那是因为主线程结束协程自动结束，主线程不会等待协程的结束
}

var (
	wait = sync.WaitGroup{}
)

func gosing() {
	fmt.Println("小林-吃饭")
	fmt.Println("小黄-吃饭")
	time.Sleep(1 * time.Second)
	fmt.Println("吃饭结束")
	//通知结束
	wait.Done()
}

// 在方法二中，我们让主线程等待协成结束。
func main1202() {
	wait.Add(4)
	go gosing()
	go gosing()
	go gosing()
	go gosing()
	wait.Wait()
	fmt.Println("主线程结束")
}

// 基本定义
func main1203() {
	//这行代码声明了一个 通道（channel），通道可以在不同的 goroutine 之间进行通信。这里的通道 c 用于传递整型数据（int 类型）。
	var c chan int //声明一个传递整形的通道。
	//初始化通道
	//缓冲通道 意味着你可以发送一个数据到通道，而不用立即等待另一个 goroutine 去读取它。这个通道最多可以存储 1 个 int 类型的数据
	//通道有两种类型：无缓冲通道 和 带缓冲通道。无缓冲通道必须要有接收者才能发送成功，否则发送操作会阻塞。带缓冲通道可以在没有接收者的情况下，存储指定缓冲区大小的数据。
	c = make(chan int, 1) //初始化一个，有一个缓冲位的通道。

	//发送数据: 这一行将 1 发送到通道 c 中。操作符 <- 表示发送数据到通道。
	//因为这个通道有一个缓冲区（大小为 1），c <- 1 不会立即阻塞，数据 1 会存储在通道的缓冲区中，等待另一个 goroutine 来读取它。
	c <- 1

	//c<-2 //会保存 deadlock  //fatal error: all goroutines are asleep - deadlock!
	//接收数据：你可以使用 val := <-c 来从通道 c 中接收数据。
	fmt.Println(<-c) //取值 1
	//fmt.Println(<-c) //在取值也会报 deadlock

	//通道的基本行为:
	//当通道 没有缓冲 时，发送和接收操作必须同时发生，否则会阻塞。
	//当通道有缓冲时，如果缓冲区没有满，发送数据时不会阻塞，直到缓冲区满为止。
	//同样，接收操作会等待，直到通道中有数据可读。

	//由于之前的通道 c 是使用 make(chan int, 1) 创建的一个 带缓冲区大小为 1 的通道，通道中现在可以存储一个值而不会阻塞。前面的已经取出。
	c <- 2
	//<-c 表示从通道 c 中接收数据，并将接收到的数据赋值给变量 n。
	//ok 是一个布尔值，用来判断接收操作是否成功。如果通道关闭且没有数据，ok 会是 false；否则为 true。
	n, ok := <-c
	//打印变量n,和ok.
	fmt.Println(n, ok)

}

func main1204() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			//通道1值的接收。
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			//通道2值的接收
			fmt.Println("received", msg2)
		}
	}
}
