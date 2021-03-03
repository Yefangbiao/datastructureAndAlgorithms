package main

import "fmt"

type cSubscriber struct {
}

func (c *cSubscriber) notify(message string) {
	fmt.Printf("get message: %v", message)
}
