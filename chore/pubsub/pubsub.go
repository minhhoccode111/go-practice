package pubsub

// problem: publish a message to multiple subscribed channel

type subscriber chan string

type publisher struct {
	subs []subscriber
}

func (p *publisher) publish(msg string) {
	for _, v := range p.subs {
		v <- msg
	}
}

func (p *publisher) subscribe() subscriber {
	sub := make(subscriber)
	p.subs = append(p.subs, sub)
	return sub
}
