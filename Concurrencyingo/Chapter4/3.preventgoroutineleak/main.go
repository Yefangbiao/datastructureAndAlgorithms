package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 谁创建，谁管理
// main goroutine 必须能够管理 它创建的goroutine
func main() {
	var wg sync.WaitGroup

	doneChannel := func() <-chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			time.Sleep(5 * time.Second)
		}()
		return done
	}

	done := doneChannel()

	newRandStream := func(done <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer wg.Done()
			defer fmt.Println("newRandStream closed")
			defer close(intStream)
			for {
				select {
				case <-done:
					return
				case intStream <- rand.Intn(5):
					time.Sleep(1 * time.Second)
				}
			}
		}()
		return intStream
	}

	consumer := func(done <-chan interface{}, intStream <-chan int) {
		defer fmt.Println("consumer closed")
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case i := <-intStream:
				fmt.Println(i)
			}
		}
	}

	wg.Add(2)
	intStream := newRandStream(done)
	go consumer(done, intStream)

	wg.Wait()

}
