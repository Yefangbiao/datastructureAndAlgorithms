package main

type shoe interface {
	GetProducer() string
	GetSize() int
	Action()
}

type iShoe struct {
	producer string
	size     int
}

func (s iShoe) GetProducer() string {
	return s.producer
}

func (s iShoe) GetSize() int {
	return s.size
}
