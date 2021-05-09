package main

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
