package main

import (
	"fmt"
	"time"
)

// 常见的for-select 循环
func main() {

	doneChannel := func() <-chan time.Time {
		return time.After(5 * time.Second)
	}
	done := doneChannel()

	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("default")
			time.Sleep(1 * time.Second)
		}
	}
}
