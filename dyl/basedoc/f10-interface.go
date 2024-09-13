// 接口
package main

import "fmt"

//接口是一组仅包含方法名,参数，返回值的未具体实现的方法的合集。
//接口本身不能绑定方法. 接口是值类型,保存的是值+原始类型。

//实现接口:
//一个类型实现了接口的所有方法,即实现了该接口。

// Animal 定义了一个 animal 的接口，它有唱，跳，rap,say的方法。
type Animal interface {
	sing()
	jump()
	rap()
	say()
}

// Chicken (鸡 需要全部实现这些接口)
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("大公鸡:", "chicken 唱了一首- 千年等一回")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳了一段 优美的舞蹈")
}
func (c Chicken) rap() {
	fmt.Println("chinck rap 说了 一段相声")
}
func (c Chicken) say() {
	fmt.Println("chinck say 请关注我,hello,hi!")
}

// Cat 需要全部实现这些接口.
type Cat struct {
	Name string
}

func (c Cat) sing() {
	fmt.Println("狸猫:", "cat 唱-喵")
}

func (c Cat) jump() {
	fmt.Println("cat 跳")
}

func (c Cat) rap() {
	fmt.Println("cat rap")
}

func (c Cat) say() {
	fmt.Println("cat say")
}

func sing(obj Animal) {

	//通过断言来获取此时的具体类型
	switch obj.(type) {
	case Chicken:
		fmt.Println("大公鸡")
	case Cat:
		fmt.Println("猫")
	}

	obj.sing()
}
func jump(obj Animal) {
	//断言某个类型。
	c, ok := obj.(Chicken) //两个参数,断言之后的类型,是否是对应类型。
	fmt.Println(c, ok)
	//如果断言报错，可以尝试获取错误。going
	//d := obj.(Cat) //一个参数 ，就是断言之后的类型,断言不对，会报错 . main.Animal is main.Chicken, not main.Cat
	//fmt.Println(d)
	obj.jump()
}

// 全部实现完成之后，chicken就是一只高级货。
func main101() {
	var animal Animal
	animal = Chicken{"战斗鸡"}
	animal.sing()
	animal.jump()
	animal.rap()
	animal.say()
}

func main102() {
	chicken := Chicken{"jk"}
	cat := Cat{"mao"}

	sing(chicken)
	jump(chicken)

	sing(cat)
	jump(cat)
}

func main() {
	//main101()
	main102()
}
