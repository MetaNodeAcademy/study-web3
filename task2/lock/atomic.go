package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	atomicCounter()
}

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func atomicCounter() {
	var num int32
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&num, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("计数器的值:", num)
}