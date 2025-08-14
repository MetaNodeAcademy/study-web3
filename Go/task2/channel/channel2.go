package main

import (
	"fmt"
	"sync"
)

func main() {
	start()
}

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/
func start() {
	count := 100
	ch := make(chan int, count)
	go producer(ch, count)
	go consumer(ch)
	time.Sleep(10 * time.Second)
}
func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到：%d\n", v)
	}
}