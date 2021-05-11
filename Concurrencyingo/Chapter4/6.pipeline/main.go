package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)

		go func() {
			defer close(intStream)
			for _, integer := range integers {
				select {
				case <-done:
					return
				case intStream <- integer:
				}
			}
		}()

		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- multiplier * i:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan interface{}, multipliedStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range multipliedStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, []int{2, 3, 4, 5}...)
	pipeline := multiply(done, add(done, multiply(done, intStream, 4), 3), 2)

	for v := range pipeline {
		fmt.Println(v)
	}

}
