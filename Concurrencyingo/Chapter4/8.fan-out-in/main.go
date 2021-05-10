package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

// 扇出,启动多个goroutine来处理来自pipeline的输入的过程
// 可选的前置条件
// 1.计算不依赖于前一个值 2.运行需要很长时间

// 扇入,将多个数据流复用合成一个流

func main() {
	done := make(chan interface{})
	defer close(done)

	numFinders := runtime.NumCPU()

	randStream := toInt(done, repeatFn(done, Rand))

	finders := make([]<-chan interface{}, numFinders)
	for i := range finders {
		finders[i] = primeFinder(done, randStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Println(prime)
	}

}

func Rand() int {
	return rand.Intn(50000)
}

func repeatFn(done <-chan interface{}, fn func() int) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for {
			select {
			case <-done:
				return
			case x := <-valueStream:
				intStream <- x.(int)
			}
		}
	}()

	return intStream
}

// 从传入的valueStream取出前num个项目
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case x := <-valueStream:
				takeStream <- x
			}
		}
	}()
	return takeStream
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// 从所有channel取值
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		defer close(multiplexedStream)
		wg.Wait()
	}()
	return multiplexedStream
}

func primeFinder(done <-chan interface{}, valueStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})

	go func() {
		defer close(primeStream)
		for {
			select {
			case <-done:
				return
			case x := <-valueStream:
				if isPrime(x) {
					primeStream <- x
				}
			}
		}
	}()

	return primeStream
}

func isPrime(x int) bool {
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
