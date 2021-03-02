package main

type caretaker struct {
	memos []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.memos = append(c.memos, m)
}

func (c *caretaker) getMemento() *memento {
	m := c.memos[len(c.memos)-1]
	c.memos = c.memos[:len(c.memos)-1]
	return m
}
