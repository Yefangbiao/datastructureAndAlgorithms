package main

type originator struct {
	state string
}

func (o *originator) State() string {
	return o.state
}

func (o *originator) SetState(state string) {
	o.state = state
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreFromMemento(m *memento) {
	o.state = m.getState()
}
