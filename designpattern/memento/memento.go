package main

type memento struct {
	state string
}

func (m *memento) getState() string {
	return m.state
}
