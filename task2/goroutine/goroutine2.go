package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建调度器实例
	scheduler := NewScheduler()

	scheduler.AddTask(func() {
		fmt.Println("任务0开始")
		time.Sleep(1 * time.Second)
		fmt.Println("任务0结束")
	})
	scheduler.AddTask(func() {
		fmt.Println("任务1开始")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务1结束")
	})
	scheduler.AddTask(func() {
		fmt.Println("任务2开始")
		time.Sleep(2 * time.Second)
		fmt.Println("任务2结束")
	})

	stats := scheduler.Run()
	// 输出统计结果
	fmt.Println("\n任务执行耗时统计:")
	for id, duration := range stats {
		fmt.Printf("任务 %d: %v\n", id, duration)
	}
}

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/
// 定义任务类型
type Task func()

// 任务执行结果结构
type TaskResult struct {
	ID       int           //任务id
	Duration time.Duration //执行耗时
}

// 任务调度器结构
type Scheduler struct {
	tasks []Task         // 任务列表
	wg    sync.WaitGroup // 等待组用于同步
}

// 创建调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]Task, 0),
	}
}

// 添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// 执行所有任务并返回执行时间统计
func (s *Scheduler) Run() map[int]time.Duration {
	count := len(s.tasks)
	if count == 0 {
		return nil
	}
	// 创建带缓冲的结果通道
	results := make(chan TaskResult, count)
	s.wg.Add(count)
	// 启动协程执行所有任务
	for id, task := range s.tasks {
		go func(taskId int, task Task) {
			defer s.wg.Done()
			start := time.Now()
			task() // 执行任务函数
			duration := time.Since(start)
			// 发送结果到通道
			results <- TaskResult{ID: id, Duration: duration}
		}(id, task)
	}
	// 等待所有任务完成
	go func() {
		s.wg.Wait()
		close(results)
	}()
	// 收集结果
	stats := make(map[int]time.Duration)
	for r := range results {
		stats[r.ID] = r.Duration
	}
	return stats
}