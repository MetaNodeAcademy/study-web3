package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	safeLock()
}

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
type Counter struct {
	mu  sync.Mutex
	num int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num++
}

func safeLock() {
	counter := Counter{num: 0}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("计数器的值:", counter.num)
}