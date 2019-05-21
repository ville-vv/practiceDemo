package main

import (
	"fmt"
	"practiceDemo/mysql_work/mysqlc"
)

func main(){
	mydb, err := mysqlc.NewMysqlClient("root","Root123", "127.0.0.1:3306", "information_schema")
	if err != nil{
		fmt.Println(err)
		return
	}
	err = mydb.Create("mysql_test_work")
	if err != nil{
		fmt.Println(err)
		return
	}
	err = mydb.Drop("mysql_test_work")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("数据库链接成功")
	mydb.Close()
}
