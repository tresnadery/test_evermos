package queue

import (
	"errors"
)

var (
	errQueueIsEmpty = errors.New("queue is empty")
)

// Queue Structure
type Queue struct {
	array []interface{}
}

// New factory to generate Queue
func New(customers ...interface{}) *Queue {
	q := &Queue{
		make([]interface{}, 0, len(customers)),
	}
	q.Enqueue(customers...)
	return q
}

// Enqueue add customer to Queue
func (q *Queue) Enqueue(customers ...interface{}) {
	q.array = append(q.array, customers...)
}

// Size return size of the Queue
func (q *Queue) Size() int {
	return len(q.array)
}

// IsEmpty check if Queue is empty or not
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

// Dequeue remove customer from Queue
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errQueueIsEmpty
	}
	customer := q.array[0]
	q.array[0] = nil
	q.array = q.array[1:]
	return customer, nil
}

// Clear clean Queue
func (q *Queue) Clear() {
	q.array = nil
}

// Peek return front of the queue
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errQueueIsEmpty
	}
	return q.array[0], nil
}

// GetCustomers retuen customers
func (q *Queue) GetCustomers() []interface{} {
	customers := make([]interface{}, 0, q.Size())
	for _, val := range q.array {
		customers = append(customers, val)
	}
	return customers
}
