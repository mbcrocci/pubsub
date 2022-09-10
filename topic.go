package pubsub

import (
	"github.com/mbcrocci/queue"
)

type Topic struct {
	topic       string
	subscribers []*subscriber
}

func (t *Topic) Subscribe(s Subscriber) {
	sub := new(subscriber)
	sub.q = queue.New[any]()
	sub.s = s
	go sub.Run()

	t.subscribers = append(t.subscribers, sub)
}

func (t Topic) BroadcastAll(data any) {
	for _, s := range t.subscribers {
		s.Publish(data)
	}
}

func (t Topic) Broadcast(s Subscriber, data any) {
	for _, sub := range t.subscribers {
		if sub.s != s {
			sub.Publish(data)
		}
	}
}
