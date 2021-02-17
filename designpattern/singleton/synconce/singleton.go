package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Once{}

type singleton struct {
}

var singletonInstance *singleton

func GetSingletonInstance() *singleton {
	if singletonInstance == nil {
		lock.Do(func() {
			fmt.Println("Creating single instance now.")
			singletonInstance = &singleton{}

		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singletonInstance
}
