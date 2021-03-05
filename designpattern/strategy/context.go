package main

type context struct {
	s strategy
}

func (c *context) SetS(s strategy) {
	c.s = s
}

func newContext() *context {
	return &context{}
}

func (c *context) executeStrategy(a, b int) {
	c.s.execute(a, b)
}
