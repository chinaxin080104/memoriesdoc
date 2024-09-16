// 线程安全与sync.map
package main

import (
	"fmt"
	"sync"
	"time"
)

var num int
var num2 int
var wait sync.WaitGroup
var lock sync.Mutex

var mp = map[string]string{}
var mp2 = sync.Map{}

func main() {
	//线程安全01-未加锁
	main1501()
	//线程安全02-加锁(同步锁)
	main1502()
	//线程安全下的 map错误
	main1503()
}
func add() {
	for i := 0; i < 1000000; i++ {
		num++
	}
	wait.Done()
}
func reduce() {
	for i := 0; i < 1000000; i++ {
		num--
	}
	wait.Done()
}

// 多线程安全，没有加锁的情况下，num的输出结果完全无法预测.
// 根本原因是 cpu的调度方式为抢占式执行,随机调度。
func main1501() {

	wait.Add(2)
	go add()
	go reduce()
	wait.Wait()
	fmt.Println(num) //0,196139,-10000 会随机输出。
}

func add2() {
	//先加锁在执行。
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num2++
	}
	//执行完循环后再解锁。
	lock.Unlock()
	wait.Done()
}
func reduce2() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num2--
	}
	lock.Unlock()
	wait.Done()
}
func main1502() {
	wait.Add(2)
	//执行逻辑是先执行完go1.add2 方法，
	go add2()
	//随后在执行go2.reduce2 方法。
	go reduce2()
	wait.Wait()
	fmt.Println(num2)
}

func reader() {
	for {
		fmt.Println(mp2.Load("time"))
	}
	wait.Done()
}

// more than one character in rune literal
func writer() {
	for {
		mp2.Store("time", time.Now().Format("15:04:05"))
	}
	wait.Done()
}

// map的线程安全错误
// 我们不能直接在并发模式下读写map.如果需要这样做，需要 1.给读写操作加锁。2.使用sync.Map
func main1503() {
	wait.Add(2)
	go reader()
	go writer()
	wait.Wait()
}
