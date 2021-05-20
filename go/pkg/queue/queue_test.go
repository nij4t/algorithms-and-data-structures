package queue

import "testing"

func TestQueueEmpty(t *testing.T) {
	queue := make(Queue, 0)

	if !queue.Empty() {
		t.Error("queue isn't empty")
	}

	queue.Enqueue(0)

	if queue.Empty() {
		t.Error("queue is empty")
	}
}

func TestDequeue(t *testing.T) {
	queue := make(Queue, 0)

	queue.Enqueue(12)

	got := queue.Dequeue()
	want := 12

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
