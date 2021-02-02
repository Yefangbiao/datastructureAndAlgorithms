package main

import "fmt"

func GetGun(name string) (iGun, error) {
	switch name {
	case "ak47":
		return newAk47(), nil
	case "m416":
		return newM416(), nil
	default:
		return nil, fmt.Errorf("getGun: ")
	}
}
