package main

// 传入一个channel,返回两个

func tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)
		for {
			select {
			case <-done:
				return
			case x := <-in:
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- x:
						out1 = nil
					case out2 <- x:
						out2 = nil
					}
				}
			}
		}
	}()

	return out1, out2
}
