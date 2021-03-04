package main

import (
	"fmt"
)

type draftState struct {
	*document
}

func (d draftState) execute() {
	fmt.Println("now is draft state, turn to moderation state")
	d.document.setState(d.moderation)
}
