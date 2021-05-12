package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	heartBeat, resultStream := DoWork(done, 500*time.Millisecond, []int{2, 3, 4, 5}...)

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

	tmpStream := make(chan int)

	go func() {
		defer close(heartBeats)
		defer close(results)

		pluse := time.Tick(pulseInterval)

		for {
			select {
			case <-done:
				return
			case <-pluse:
				select {
				case heartBeats <- struct{}{}:
				default:
				}

			case x, ok := <-tmpStream:
				if !ok {
					return
				}
				results <- x
			}
		}

	}()

	// 模拟处理函数
	work := func() {
		defer close(tmpStream)
		i := 0
		for {
			select {
			case <-done:
				return
			case tmpStream <- nums[i] * 2:
				i++
				if i >= len(nums) {
					return
				}
				time.Sleep(2 * time.Second)
			}
		}
	}
	go work()

	return heartBeats, results
}
