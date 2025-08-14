package Sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
var db *sqlx.DB

// 连接数据库
func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:Xm@123@tcp(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
}

type Employees struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func GetEmpListByDeptName() []Employees {
	var emp []Employees
	err := db.Select(&emp, "SELECT id ,name,department,salary FROM employees WHERE department=?", "技术部")
	if err != nil {
		fmt.Printf("select employees failed, err:%v\n", err)
		return nil
	}
	return emp
}
func GetMaxSalaryEmp() Employees {
	var emp Employees
	err := db.Get(&emp, "SELECT id ,name,department,salary FROM employees order by salary desc limit 1")
	if err != nil {
		fmt.Printf("select employees failed, err:%v\n", err)
		return emp
	}
	return emp
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type Book struct {
	Id     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func GetBooks() []Book {
	var books []Book
	err := db.Select(&books, "SELECT id,title,author,price FROM books WHERE price > ?", 50)
	if err != nil {
		fmt.Printf("select books failed, err:%v\n", err)
	}
	return books
}
