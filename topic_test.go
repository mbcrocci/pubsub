package pubsub

import (
	"testing"
	"time"
)

type TestSubscriber struct {
	messages int
}

func (t *TestSubscriber) OnMessage(data any) {
	t.messages++
}

func TestTopic(t *testing.T) {
	topic := new(Topic)
	s1 := new(TestSubscriber)
	s2 := new(TestSubscriber)

	topic.Subscribe(s1)
	topic.Subscribe(s2)

	topic.BroadcastAll(1)

	// need to sleep to allow subscriber go routines to have time to dequeue
	time.Sleep(10 * time.Millisecond)

	if s1.messages != 1 || s2.messages != 1 {
		t.FailNow()
	}

	topic.Broadcast(s2, 1)
	time.Sleep(10 * time.Millisecond)

	if s1.messages != 2 || s2.messages != 1 {
		t.FailNow()
	}
}
