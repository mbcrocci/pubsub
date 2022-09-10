package pubsub

import (
	"testing"
	"time"
)

type TestPubSubSubscriber struct {
	messages int
}

func (t *TestPubSubSubscriber) OnMessage(data any) {
	t.messages++
}

func TestPubSub(t *testing.T) {
	ps := NewPubSub()

	s1 := new(TestPubSubSubscriber)
	s2 := new(TestPubSubSubscriber)
	s3 := new(TestPubSubSubscriber)

	ps.Subscribe(s1, "first topic")
	ps.Subscribe(s2, "first topic")
	ps.Subscribe(s3, "second topic")

	if len(ps.topics) != 2 {
		t.FailNow()
	}

	ps.BroadcastAll("first topic", 1)
	ps.BroadcastAll("first topic", 1)

	if (s1.messages != 2 || s2.messages != 2) && s3.messages != 0 {
		t.FailNow()
	}

	ps.BroadcastAll("second topic", 1)
	time.Sleep(10 * time.Millisecond)

	if (s1.messages != 2 || s2.messages != 2) && s3.messages != 1 {
		t.FailNow()
	}

	ps.Broadcast(s1, "first topic", 1)
	time.Sleep(10 * time.Millisecond)

	if (s1.messages != 2 || s2.messages != 3) && s3.messages != 1 {
		t.FailNow()
	}
}

func BenchmarkPubSub(b *testing.B) {
	ps := NewPubSub()

	s1 := new(TestPubSubSubscriber)
	s2 := new(TestPubSubSubscriber)
	ps.Subscribe(s1, "topic")
	ps.Subscribe(s2, "topic")

	for i := 0; i < b.N; i++ {
		ps.Broadcast(s1, "topic", i)
	}

}
