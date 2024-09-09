// 数组，切片，map
package main

import (
	"fmt"
	"sort"
)

func main() {
	//func array demo
	//区别
	// 数组：[4]string{"a", "b", "c", "d"} 明确指定数组长度为 4。
	// 切片：[]string{"a", "b", "c", "d"} 没有明确长度，表示一个切片，长度可以动态变化。
	//arrayfunc()
	//arrayslice()
	// funcmake()
	// slices()
	funcmap()
}

// 数组
func arrayfunc() {
	// 数组里面的元素必须全部为同一类型。
	// 声明数组时，必须指定其长度或者大小 。

	//eg:1
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)

	//eg:2
	var array1 = [3]int{4, 5, 6}
	fmt.Println(array1)
	//eg:3  在 Go 语言中，[...] 是一种特殊的语法，用于在数组声明时让编译器根据初始化值的数量自动推导数组的长度。
	var array2 = [...]int{7, 8, 9}
	fmt.Println(array2)

	//索引-如果需要修改某个值，只能根据索引去找然后替换。
	array2[2] = 10
	fmt.Println(array2)

	//正向索引就是从0开始。从左往右。
	//负向索引
	//是的，var sList = []string{"a", "b", "c", "d"} 这样的语法是完全可以的，它定义了一个 切片（slice），而不是一个数组。
	//在 Go 中，切片是一个动态数组，它的长度可以在运行时自动调整，和数组不同的是，切片的长度是可变的，而数组的长度是固定的。
	var sList = []string{"a", "b", "c", "d"}
	// 如果在go语言里面，我想拿到倒数第二个元素，怎么办？直接减2就可以了.
	fmt.Println(sList[len(sList)-2])

}

// 切片
// 很明显啊，go里面的数组，长度被限制死了，所以不经常用 ,所以go出了一个数组plus，叫做slice（切片）
func arrayslice() {
	//切片相对于数组更灵活，因为在声明切片后其长度是可变的。
	var list []string

	//切片中加值
	list = append(list, "小小")
	list = append(list, "木木")
	fmt.Println(list)
	fmt.Println(len(list))

	list[1] = "万能的您,赶紧干正事吧"
	fmt.Println(list)
}

// make函数
func funcmake() {
	//除了基本数据类型，其他数据如果只定义不赋值，那么实际的值就是nil
	var list []string
	fmt.Println(list == nil) //true
	//那么我们可以通过make函数创建指定长度,指定容量的切片了。
	//make([]type,length,capacity)
	//make([]string, 0)：创建了一个长度为 0、容量也为 0 的字符串切片 string2。make 函数用于分配和初始化切片。
	//长度（len(string2)）：切片当前包含的元素个数，0。
	//容量（cap(string2)）：切片底层数组最多能够容纳的元素个数，初始化时为 0。

	var string2 = make([]string, 0)
	fmt.Println(string2, len(string2), cap(string2))
	fmt.Println(string2 == nil) //false

	//切片
	list2 := make([]int, 2, 2)
	fmt.Println(list2, len(list2), cap(list2))
	fmt.Println(list2[1])
}

// 切片是数组切出来的
func slices() {
	var list1 = [...]string{"a", "b", "c"}
	//[:] 是 Go 语言中常用的切片表达式，可以表示从头到尾整个数组或者切片。
	// list[:]：这是特殊的操作，表示从头开始切到尾。相当于对数组从索引 0 开始到数组末尾进行切片，这样生成的 slices 就是原数组的整个部分，只是以切片的形式表示。
	slices := list1[:]
	fmt.Println(slices)

	// 左一刀，右一刀  变成了切片
	// array[start:end] 这种形式表示对数组（或切片）进行切片操作，从索引 start 开始（包含 start 位置），到索引 end 结束（不包含 end 位置）。
	// 左一刀：表示你从数组的某个索引开始（切片操作的起点），也就是左侧的起始索引。例如 list[1:]，你从索引 1 开始切片。
	// 右一刀：表示你从数组的某个索引结束（切片操作的终点），也就是右侧的结束索引。例如 list[:2]，你从头开始切到索引 2（不包括 2）。
	fmt.Println(list1[1:2]) // b

	// 切片排序
	var list = []int{4, 5, 3, 2, 7}
	fmt.Println("排序前:", list)
	sort.Ints(list)
	fmt.Println("升序:", list)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println("降序:", list)
}

// map是一种内置的数据结构，他是一个无序的key-value集合
func funcmap() {
	//声明
	var m1 map[string]string
	//初始化1
	//m1 = make(map[string]string)
	//初始化2
	m1 = map[string]string{}

	//eg 声明并赋值
	var a1 = map[string]string{}
	var a2 = make(map[string]string)

	//设置值
	m1["name"] = "董养龙"
	a1["az"] = "abcd-----z"
	a2["AZ"] = "ABCDEF----Z"
	fmt.Println(m1["name"])
	delete(m1, "name")
	fmt.Println(m1)
	//测试map 用法
	fmt.Println(a1, a2)
}
