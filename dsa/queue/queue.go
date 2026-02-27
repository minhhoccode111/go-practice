package queue

type IQueue[T any] interface {
	IsEmpty() bool
	Enqueue(T)
	Dequeue() T
}

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	f := q.items[0]
	q.items = q.items[1:]
	return f
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}
