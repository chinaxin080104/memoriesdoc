// 判断语句
package main

import "fmt"

func main() {
	//if语句
	//iffunc()
	//switch语句
	switchfunc()
}

func iffunc() {
	fmt.Println("请输入您的绩效分数:")
	//定义一个值获取积分
	//  <=60  不及格
	//  60-75  在努力下
	//  75-95  加油
	//  95 已上，加油
	var jifen int
	fmt.Scan(&jifen)
	// 嵌套执行
	if jifen <= 60 {
		fmt.Println("不及格")
	} else {
		//中断式实现，它也有个好听的名字，叫卫语句
		if jifen >= 95 && jifen <= 120 {
			fmt.Println("优秀")
		}
		if jifen < 95 && jifen >= 75 {
			fmt.Println("加油哦")
		}
		//多条件判断 &&  || 逻辑运算符
		if jifen < 75 && jifen > 60 {
			fmt.Println("在努力一下")
		}
		if jifen > 120 {
			fmt.Println("飘了,嘿嘿！")
		}
	}

}

func switchfunc() {
	//switch 语句特点,直观。
	//eg:根据分数，给出是否是好学生。
	fmt.Println("请输入您的英语考试成绩:")
	var cj int
	fmt.Scan(&cj)
	switch {
	case cj <= 0:
		fmt.Println("好好学习去!")
	case cj <= 30:
		fmt.Println("您的老师是数学老师吗?")
	case cj <= 61:
		fmt.Println("加油，努力记单词!")
	case cj <= 80:
		fmt.Println("马马虎虎!")
	default:
		fmt.Println("您一定是个好学生，好好学习!")
	}

	//判断输入的天数是那天
	fmt.Println("请输入今天的星期-已数字输入:")
	var weekday int
	fmt.Scan(&weekday)

	switch weekday {
	case 1:
		fmt.Println("今天是周一")
	case 2:
		fmt.Println("今天是周二")
	case 3:
		fmt.Println("今天是周三")
	case 4:
		fmt.Println("今天是周四")
	case 5:
		fmt.Println("今天是周五")
	case 6, 7:
		fmt.Println("今天是周末,有一起出去玩的吗?")
	default:
		fmt.Println("小伙子,请注意听题!")
	}

}
