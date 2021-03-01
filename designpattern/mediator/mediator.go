package main

type mediator interface {
	canArrive(train) bool
	notifyDepart()
}
