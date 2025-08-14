package main

import (
	"fmt"
)

func main() {
	//指针1
	// a := 10
	// pointer1(&a)
	// fmt.Println(a)

	//指针2
	array := &[]int{1, 2, 3, 4}
	pointer2(array)
	fmt.Println(*array)
}

/*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func pointer2(array *[]int) {
	if array == nil || *array == nil {
		return
	}
	for i, v := range *array {
		(*array)[i] = v * 2
	}
}

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
*/
func pointer1(a *int) {
	*a += 10
}