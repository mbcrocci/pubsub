package pubsub

import "github.com/mbcrocci/queue"

type subscriber struct {
	s Subscriber
	q *queue.Queue[any]
}

func (s *subscriber) Publish(data any) {
	s.q.Enqueue(&data)
}

func (s *subscriber) ProcessMessage() {
	data := s.q.Dequeue()
	s.s.OnMessage(&data)
}

func (s *subscriber) Run() {
	go func() {
		for {
			s.ProcessMessage()
		}
	}()
}
