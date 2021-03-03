package main

type subscriber interface {
	notify(message string)
}
