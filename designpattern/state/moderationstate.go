package main

import "fmt"

type moderationState struct {
	*document
}

func (d moderationState) execute() {
	fmt.Println("now is moderation state, turn to publish state")
	d.document.setState(d.publish)
}
