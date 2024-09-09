// 循环体基础
package main

import (
	"fmt"
	"time"
)

func main() {
	//任何编程都有for循环，它的一般写法是 for 初始条件;条件;操作{}
	//例如求和
	funfor()

	//for循环的5种变体。
	//funfor01()
	funfor02()
	funfor03()
	funfor04()
	funfor05()
}

// 正常的for循环
func funfor() {
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 死循环
func funfor01() {
	for {
		time.Sleep(1 * time.Second)
		//输出 2024-09-07 17:34:23
		//Go语言中的time.Format使用了一种非常独特的时间格式化方式。它通过固定的参考时间 2006-01-02 15:04:05 来定义日期和时间的格式。这组数字并不是随机的，而是一个特定的日期和时间：
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}
}

// while模式
func funfor02() {
	// 由于go没有while循环,如果需要，则是有for循环变化产生。
	i := 0
	sum := 0
	//添加上<=后值 就跟funfor中的值一样了
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

// do-while模式 ;先执行一次，在判断
func funfor03() {
	i := 0
	sum := 0
	for {
		sum += i
		i++
		//不满足条件时，跳出
		if i == 101 {
			//break用于跳出当前循环
			break
		}
	}
	fmt.Println(sum)
}

// 遍历切片,map研究。
func funfor04() {
	// 第 1 首歌曲是 年轮
	// 第 2 首歌曲是 光亮
	// 第 3 首歌曲是 追光者
	// 第 4 首歌曲是 勇敢的心
	// 第 5 首歌曲是 走南闯北
	// 第 6 首歌曲是 千年缘
	// 第 7 首歌曲是 成都

	//遍历切片
	sing := []string{"年轮", "光亮", "追光者", "勇敢的心", "走南闯北", "千年缘", "成都"}
	//range 是一个函数。赋值给了index,s2 。   for循环-循环的是sing这个切片。
	for index, s2 := range sing {
		key := index + 1 //index 相当于forearch 里面的key,从0起步，所有加+1后，就是歌曲ID

		//输出数值。
		fmt.Println("第", key, "首歌曲是", s2)
	}

	//遍历map
	list := map[string]int{
		"id":     1,
		"age":    22,
		"price":  100,
		"statue": 0,
	}
	//这里range 把list -map里面的key-value 值遍历循环打印出来。
	for key, val := range list {
		fmt.Println(key, val)

		//变量 打印map里面的 （值--键）。
		fmt.Println(list[key], key)
	}

}

// 循环的终止与跳出
// break用于跳出当前循环
// contionue 用于跳过本轮循环
func funfor05() {
	//9*9 乘法表
	// 1 * 1=1
	// 2 * 1=2         2 * 2=4
	// 3 * 1=3         3 * 2=6         3 * 3=9
	// 4 * 1=4         4 * 2=8         4 * 3=12        4 * 4=16
	// 5 * 1=5         5 * 2=10        5 * 3=15        5 * 4=20        5 * 5=25
	// 6 * 1=6         6 * 2=12        6 * 3=18        6 * 4=24        6 * 5=30        6 * 6=36
	// 7 * 1=7         7 * 2=14        7 * 3=21        7 * 4=28        7 * 5=35        7 * 6=42        7 * 7=49
	// 8 * 1=8         8 * 2=16        8 * 3=24        8 * 4=32        8 * 5=40        8 * 6=48        8 * 7=56        8 * 8=64
	// 9 * 1=9         9 * 2=18        9 * 3=27        9 * 4=36        9 * 5=45        9 * 6=54        9 * 7=63        9 * 8=72     9 * 9=81
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			m := i * j
			fmt.Printf("%d * %d=%d \t", i, j, m)
		}
		fmt.Printf("\n")
	}

	//另外一种写法

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {

				//continue 用于跳过本次循环。
				continue
			}

			m := i * j
			fmt.Printf("%d*%d=%d \t", i, j, m)
		}
		fmt.Println()
	}

}
