package polling

import "sync"

type Queue[T any] struct {
	items    []T
	lock     *sync.RWMutex
	capacity int
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		items:    make([]T, 0, capacity),
		lock:     new(sync.RWMutex),
		capacity: capacity,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// If queue is at capacity, dequeue oldest item
	if len(q.items) == q.capacity {
		q.items = q.items[1:]
	}

	// Enqueue new item
	q.items = append(q.items, item)
}

// Return a copy of q.items
func (q *Queue[T]) Copy() []T {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return append([]T(nil), q.items...)
}
