package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutine1()
}

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/
func goroutine1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done() // 确保协程结束时通知 WaitGroup
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数：", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数：", i)
		}
	}()
	wg.Wait()
}