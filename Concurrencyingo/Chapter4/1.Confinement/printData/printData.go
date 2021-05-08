package main

import (
	"bytes"
	"fmt"
	"sync"
)

// 不会发生改变的数据 or 受到保护的数据 是隐式并发安全的
func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buf bytes.Buffer
		for _, b := range data {
			buf.WriteByte(b)
		}
		fmt.Printf("Received:%s\n", buf.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("Golang")

	printData(&wg, data[:3])
	printData(&wg, data[3:])

	wg.Wait()
}
