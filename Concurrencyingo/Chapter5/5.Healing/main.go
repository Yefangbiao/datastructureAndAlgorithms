package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 可以监控和重启的goroutine信号
type startGoroutineFn func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (heartbeat <-chan interface{})

func newSteward(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
	// 一个管理员监控goroutine需要的timeout变量，还有一个函数startGoroutine来启动它监控的goroutine
	return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})

		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() {
				// 定义了一个闭包,提供了一个统一的方法来启动我们正在监视的goroutine
				wardDone = make(chan interface{}) // 创建了一个停止信号
				// 启动要监视的goroutine.
				// 1.wardDone 停止这个startGoroutine函数
				// 2.done 管理员也停止了
				wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
			}
			startWard()
			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				// 确保管理员发出心跳
				timeoutSignal := time.After(timeout)

				for {
					select {
					case <-done:
						return
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						// 如果收到了goroutine的心跳，继续执行监控循环
						continue monitorLoop
					case <-timeoutSignal:
						// 如果在暂停期间没有收到管理区startGoroutine goroutine的心跳,要求goroutine停止
						// 并且重新开启
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					}
				}

			}
		}()

		return heartbeat
	}
}

func doWorkFn(done <-chan interface{}, intList ...int) (startGoroutineFn, <-chan interface{}) {
	intChanStream := make(chan (<-chan interface{}))
	// 内存泄漏了,应该是分配到了堆上
	intStream := bridge(done, intChanStream)

	doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		intStream := make(chan interface{})
		heartbeat := make(chan interface{})
		go func() {
			defer close(intStream)
			select {
			case <-done:
				return
			case intChanStream <- intStream:
				// 这里将新的intStream 加入桥接intStream
			}

			pulse := time.Tick(pulseInterval)

			for {
			valueLoop:
				for _, intVal := range intList {
					if intVal < 0 {
						log.Printf("negative value: %v\n", intVal)
					}
					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case intStream <- intVal:
							continue valueLoop
						case <-done:
							return
						}
					}
				}
			}
		}()
		return heartbeat
	}
	return doWork, intStream
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)

	for intVal := range take(done, intStream, 6) {
		fmt.Printf("Receive :%v\n", intVal)
	}
}
