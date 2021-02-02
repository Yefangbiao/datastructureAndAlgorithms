package main

import "fmt"

func main() {
	ak, err := GetGun("ak47")
	if err != nil {
		panic(err)
	}
	m416, err := GetGun("m416")
	if err != nil {
		panic(err)
	}
	printGun(ak)
	printGun(m416)
}

func printGun(g iGun) {
	fmt.Println("name: ", g.getName(), "power: ", g.getPower())
}
