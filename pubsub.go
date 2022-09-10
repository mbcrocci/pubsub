package pubsub

type PubSub struct {
	topics map[string]*Topic
}

func NewPubSub() *PubSub {
	ps := new(PubSub)
	ps.topics = make(map[string]*Topic)

	return ps
}

func (p *PubSub) Subscribe(s Subscriber, topic string) {
	if _, exists := p.topics[topic]; exists {
		p.topics[topic].Subscribe(s)
	} else {
		t := new(Topic)
		t.topic = topic
		t.subscribers = make([]*subscriber, 0)
		t.Subscribe(s)

		p.topics[topic] = t
	}
}

func (p *PubSub) BroadcastAll(topic string, data interface{}) {
	if topic, exists := p.topics[topic]; exists {
		topic.BroadcastAll(data)
	}
}

func (p *PubSub) Broadcast(s Subscriber, topic string, data interface{}) {
	if topic, exists := p.topics[topic]; exists {
		topic.Broadcast(s, data)
	}
}

type Subscriber interface {
	OnMessage(data any)
}
