package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {
}

var singletonInstance *singleton

func GetSingletonInstance() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creating single instance now.")
			singletonInstance = &singleton{}
			return singletonInstance
		} else {
			fmt.Println("Creating single is already created.")
			return singletonInstance
		}
	}
	fmt.Println("Creating single is already created.")
	return singletonInstance
}
