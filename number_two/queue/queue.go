package queue

import(
	"errors"
)
var(
	errQueueIsEmpty = errors.New("queue is empty")
)
// Queue Structure
type Queue struct{
	array []interface{}
}

// New factory to generate Queue
func New(customers ...interface{})*Queue{
	q := &Queue{
		make([]interface{}, 0, len(customers)),
	}
	q.Enqueue(customers...)
	return q
}
// Enqueue add customer to Queue
func (q *Queue) Enqueue(customers ...interface{}){
	q.array = append(q.array, customers...)
}