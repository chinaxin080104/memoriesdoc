// 结构体
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//结构体定义
	//main01()
	//继承
	//main02()
	//结构体指针
	//main03()
	//结构图指针方法
	//main04()
	//结构体tag
	//main05()
	//json tag
	main06()
}

//结构体定义
//type 结构图名称 struct {
//	名称 类型
//}

// 定义一个学生结构体
type Students struct {
	Id   int
	Name string
	Age  int
}

// 给学校送一个学生
func (s *Students) PrintInfo() string {
	return fmt.Sprintf("id:%d,name:%s,age:%d", s.Id, s.Name, s.Age)
}

func main01() {
	s := Students{
		Id:   1,
		Name: "乾隆",
		Age:  24,
	}

	s.Name = "雍正"
	s.Age = 31
	s.Id = 2

	m := s.PrintInfo()
	fmt.Println(m)
}

// --------------------------------------------------------------------------------------------
// 结构体的继承
type People struct {
	Time string
}

func (p People) Info() {
	fmt.Println("people", p.Time)
}

type Student2 struct {
	People
	Name string
	Age  int
}

func (s Student2) PrintInfo() {
	fmt.Printf("name:%s age:%d\n", s.Name, s.Age)
}
func main02() {
	p := People{
		Time: "2024-10-15 14:51",
	}
	s := Student2{
		People: p,
		Name:   "关羽",
		Age:    21,
	}
	s.Name = "周先生"
	s.PrintInfo()
	s.Info() //调用父结构体方法

	//time是一个时间
	fmt.Println(s.People.Time)
	fmt.Println(s.Time)

}

// -------------------------------------------------------
// 结构体指针  根据值传递和引用传递。如果我们需要在函数里面或者方法里面修改结构体里面的属性。只能使用结构体指针或者指针方法。
type Student struct {
	Name string
	Age  int
}

func SetAge(info Student, age int) {
	info.Age = age
}
func SetAge1(info *Student, age int) {
	info.Age = age
}

func main03() {
	s := Student{
		Name: "小林",
		Age:  24,
	}
	fmt.Println(s.Age)
	SetAge(s, 18) //值是替换不了 24的值的。
	fmt.Println(s.Age)

	SetAge1(&s, 19)
	fmt.Println(s.Age)
	//执行结果如下
	//24
	//24
	//19
}

// 指针方法
func (s Student) SetAge(age int) {
	s.Age = age
}

func (s *Student) SetAge1(age int) {
	s.Age = age
}

func main04() {
	s := Student{
		Name: "小曹",
		Age:  21,
	}
	s.SetAge(18) //返回21没有更新
	fmt.Println(s.Age)
	s.SetAge1(19) //返回19值，更新了
	fmt.Println(s.Age)
}

type Book struct {
	Name string `json:"name"`
	Page int    `json:"page"`
}

func main05() {
	s1 := Book{
		Name: "四大名著之西游记",
		Page: 1832,
	}
	s2 := Book{
		Name: "四大名著之三国演义",
		Page: 2812,
	}
	s3 := Book{
		Name: "四大名著之水浒传",
		Page: 861,
	}
	s4 := Book{
		Name: "四大名著之红楼梦",
		Page: 17832,
	}
	byteData1, _ := json.Marshal(s1) //-是不接收。
	byteData2, _ := json.Marshal(s2)
	byteData3, _ := json.Marshal(s3)
	byteData4, _ := json.Marshal(s4)
	fmt.Println(string(byteData1))
	fmt.Println(string(byteData2))
	fmt.Println(string(byteData3))
	fmt.Println(string(byteData4))
}

// omitempty 空值省略
type Car struct {
	Name    string `json:"name"`
	Price   int    `json:"price,omitempty"`
	Company string `json:"company"`
}

func main06() {
	Cars := Car{
		Name:    "吉利",
		Price:   12,
		Company: "浙江省吉利",
	}
	Cars2 := Car{
		Name:    "宝马",
		Price:   0,
		Company: "苏州宝马",
	}
	byteData, _ := json.Marshal(Cars)
	byteData2, _ := json.Marshal(Cars2)
	fmt.Println(string(byteData), string(byteData2))
}
