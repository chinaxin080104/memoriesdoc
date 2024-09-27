package main

import (
	"fmt"
	"gorm.io/gorm"
)

func main() {
	username := "root"
	password := "root"
	host := "127.0.0.1"   //数据库地址
	port := 3306          //数据库端口
	Dbname := "zero_user" //数据库名
	timeout := "10s"      //链接超时10s.
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, charset, timeout)
	//链接数据库
	mysql := mysql.Open(dsn)
	db, err := gorm.Open(mysql)
	//if err != nil {
	//	fmt.Println(err)
	//}
	////如果成功的话,提示
	//fmt.Println(db)
}
