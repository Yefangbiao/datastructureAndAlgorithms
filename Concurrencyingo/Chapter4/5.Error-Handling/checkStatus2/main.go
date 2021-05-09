package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	err  error
	resp *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {

				resp, err := http.Get(url)

				res := Result{
					err:  err,
					resp: resp,
				}

				select {
				case <-done:
					return
				case results <- res:
				}

			}

		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://www.a"}
	results := checkStatus(done, urls...)
	for res := range results {
		if res.err != nil {
			fmt.Printf("err: %v\n", res.err)
			continue
		}
		fmt.Println(res.resp.Status)
	}
}
