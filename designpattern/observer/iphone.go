package main

type iphone struct {
	p *publisher
}

func (i *iphone) newMessage(message string) {
	i.p.notifyAll(message)
}
