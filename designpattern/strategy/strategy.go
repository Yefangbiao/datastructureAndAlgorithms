package main

type strategy interface {
	execute(a, b int)
}
