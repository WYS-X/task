package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type safeCounter struct {
	mu    sync.Mutex
	count int
}

func (n *safeCounter) Increment() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.count++
}
func (n *safeCounter) GetCount() int {
	n.mu.Lock()
	return n.count
}
func q1() {
	var wg sync.WaitGroup
	c := safeCounter{count: 0}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				c.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("锁最终计数：", c.GetCount())
}
func q2() {
	var counter atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("原子操作最终计数：", counter.Load())
}
func main() {
	//使用锁
	q1()

	//原子操作
	q2()
}
