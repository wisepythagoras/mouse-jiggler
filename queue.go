package main

type Queue struct {
	MaxSize int
	queue   [][]int
}

// Push adds to the queue.
func (q *Queue) Add(item []int) {
	if len(q.queue) > q.MaxSize {
		// Remove last item from queue.
		q.queue = q.queue[1:]
	}

	q.queue = append(q.queue, item)
}

// Gets an item from the queue.
func (q *Queue) Pop() []int {
	if len(q.queue) == 0 {
		return nil
	}

	item := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]

	return item
}

// Get returns the raw queue.
func (q *Queue) Get() [][]int {
	return q.queue
}
