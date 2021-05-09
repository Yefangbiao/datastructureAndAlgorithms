package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {

				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}

			}

		}()

		return responses
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://www.bad.com"}
	responses := checkStatus(done, urls...)
	for response := range responses {
		fmt.Println(response.Status)
	}

}
