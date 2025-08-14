package main

import (
	"fmt"
)

func main() {
	employee := Employee{
		Person: Person{
			Name: "战三", Age: 1000,
		},
		EmployeeID: 1,
	}
	employee.PrintInfo()
}

/*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
type Person struct {
	Name string
	Age  uint
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("EmployeeID:", e.EmployeeID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Person.Name:", e.Person.Name)
	fmt.Println("Person.Age:", e.Person.Age)
}