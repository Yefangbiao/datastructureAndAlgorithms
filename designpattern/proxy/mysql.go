package main

import "fmt"

type mysql struct {
}

func (m *mysql) connect() {
	fmt.Println("connected is success!")
}
