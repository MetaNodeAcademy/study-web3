package main

import (
	"encoding/json"
	"fmt"
	"task3/gormtest"
)

func main() {
	/*
		gorm
	*/
	//初始化表结构
	//gormtest.InitTables()
	//初始化数据
	//gormtest.InitData()

	//user := gormtest.GetPostByUserId(1)
	//data, _ := json.Marshal(user)
	//fmt.Println("查询某个用户发布的所有文章及其对应的评论信息: ", string(data))

	post := gormtest.GetPostWithMostComments()
	data, _ := json.Marshal(post)
	fmt.Println("查询评论数量最多的文章信息: ", string(data))

	/*
		Sqlx
	*/
	//emps := Sqlx.GetEmpListByDeptName()
	//fmt.Println("查询所有部门为\"技术部\"的员工信息: ", emps)
	//emp := Sqlx.GetMaxSalaryEmp()
	//fmt.Println("查询工资最高的员工信息: ", emp)
	//books := Sqlx.GetBooks()
	//fmt.Println("查询价格大于50元的书籍: ", books)
}
