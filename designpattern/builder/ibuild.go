package main

type iBuild interface {
	setWall()
	setWindow()
	setDoor()
	getResult() house
}
