package main

type publisher struct {
	subscribers []subscriber
}

func (p *publisher) addSubscriber(sub subscriber) {
	p.subscribers = append(p.subscribers, sub)
}

func (p *publisher) deleteSubscriber(sub subscriber) {
	for i, s := range p.subscribers {
		if s == sub {
			rear := p.subscribers[i+1]
			p.subscribers = p.subscribers[:i]
			p.subscribers = append(p.subscribers, rear)
			return
		}
	}
}

func (p *publisher) notifyAll(message string) {
	for _, s := range p.subscribers {
		s.notify(message)
	}
}
