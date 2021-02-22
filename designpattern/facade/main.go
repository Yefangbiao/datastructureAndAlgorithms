package main

import (
	"fmt"
	"log"
)

func main() {
	newA := NewWalletFacade("abc", "123")

	if err := newA.AddAmount("abc", "123", 1423); err != nil {
		log.Print(err)
	}

	if err := newA.ReduceAmount("abc", "123", 234); err != nil {
		log.Print(err)
	}

	fmt.Println(newA.balance)
}
