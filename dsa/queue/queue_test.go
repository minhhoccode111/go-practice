package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("New queue is empty", func(t *testing.T) {
		q := New[int]()
		if !q.IsEmpty() {
			t.Error("Expected empty queue")
		}
	})

	t.Run("Enqueue make queue non-empty", func(t *testing.T) {
		q := New[int]()
		q.Enqueue(1)
		if q.IsEmpty() {
			t.Error("Expected non-empty queue")
		}
	})

	t.Run("Dequeue returns items in FIFO order", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		q := New[string]()

		for _, v := range s {
			q.Enqueue(v)
		}

		for _, want := range s {
			if got := q.Dequeue(); got != want {
				t.Error("")
			}
		}
	})

	t.Run("Dequeue on empty panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic but got none")
			}
		}()
		q := New[int]()
		q.Dequeue()
	})

	t.Run("Works with custom type", func(t *testing.T) {
		type custom struct{ a, b int }
		q := New[custom]()
		q.Enqueue(custom{1, 2})
		q.Enqueue(custom{3, 4})
		if got := q.Dequeue(); got != (custom{1, 2}) {
			t.Errorf("Unexpected value %v", got)
		}
	})
}
