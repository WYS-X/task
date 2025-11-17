package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//打印奇偶数
	q1()

	//并发执行任务
	execute([]Task{do1, do2, do3, do4})
}

type Task func()

func do1() {
	time.Sleep(time.Millisecond * 500)
}
func do2() {
	time.Sleep(time.Millisecond * 600)
}
func do3() {
	time.Sleep(time.Millisecond * 700)
}
func do4() {
	time.Sleep(time.Millisecond * 800)
}
func execute(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for i, t := range tasks {
		go func() {
			defer wg.Done()
			start := time.Now()
			t()
			fmt.Println("第", i+1, "个任务完成，用时：", time.Since(start))
		}()
	}
	wg.Wait()
}

func q1() {
	var wg sync.WaitGroup
	wg.Add(2)
	//打印奇数，偶数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 > 0 {
				fmt.Println("协程1奇数：", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("协程2偶数：", i)
			}
		}
	}()
	wg.Wait()
}
