package main

import "fmt"

type account struct {
	name string
}

func (a *account) Check(checkName string) error {
	if checkName != a.name {
		return fmt.Errorf("error: acount name is wrong\n")
	}
	return nil
}
