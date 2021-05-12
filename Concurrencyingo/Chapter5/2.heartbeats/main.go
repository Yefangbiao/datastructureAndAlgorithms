package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	heartBeat, resultStream := DoWork(done, 200*time.Millisecond, []int{2, 3, 4, 5}...)

	for {
		select {
		case <-heartBeat:
			fmt.Println("收到心跳")
		case x, ok := <-resultStream:
			if !ok {
				return
			}
			fmt.Printf("收到结果: %v\n", x)
		case <-time.After(3 * time.Second):
			return
		}
	}
}

// 模拟处理，处理每一个num可能时间很长,通过心跳让客户端知晓还在处理
func DoWork(done <-chan interface{}, pulseInterval time.Duration, nums ...int) (<-chan interface{}, <-chan int) {
	heartBeats := make(chan interface{}, 1)
	results := make(chan int)

	// 心跳用来确保还在执行
	go func() {

		pulse := time.Tick(pulseInterval)
		for {
			select {
			case <-done:
				return
			case <-pulse:
				select {
				case heartBeats <- struct{}{}:
				default:
				}
			}
		}
	}()

	// 模拟处理函数
	work := func() {
		defer close(results)
		defer close(heartBeats)

		for _, num := range nums {
			select {
			case <-done:
				return
			case results <- num * 2:
				time.Sleep(2 * time.Second)
			}
		}
	}
	go work()

	return heartBeats, results
}
