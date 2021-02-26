package main

type handle interface {
	setNext(next department)
	execute(p *patient)
}
