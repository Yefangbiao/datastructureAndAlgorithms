package main

type m416 struct {
	gun
}

func newM416() *m416 {
	return &m416{gun: gun{
		name:  "m416",
		power: 38,
	}}
}
