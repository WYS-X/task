package main

import (
	"fmt"
	"sync"
)

func q1() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("发送", i)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			r := <-ch
			fmt.Println("接收到：", r)
		}
	}()
	wg.Wait()
}
func q2() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Println("发送数据：", i)
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for m := range ch {
			fmt.Println("接收到", m)
		}
	}()
	wg.Wait()
}
func main() {
	//协程间通信
	q1()

	//带有缓冲的通道
	q2()
}
