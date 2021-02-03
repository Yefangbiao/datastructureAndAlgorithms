package main

type shirt interface {
	GetProducer() string
	GetSize() int
	Action()
}

type iShirt struct {
	producer string
	size     int
}

func (s iShirt) GetProducer() string {
	return s.producer
}

func (s iShirt) GetSize() int {
	return s.size
}
