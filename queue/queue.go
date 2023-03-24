package queue

import (
	"fmt"
)

type emptyQueueError struct {
}

func (e emptyQueueError) Error() string {
	return fmt.Sprintf("queue is empty")
}

type Queue struct {
	arr []any
}

func MakeQueue() Queue {
	return Queue{make([]any, 0)}
}

func (q *Queue) returnQueue() []any {
	return q.arr
}

func (q *Queue) Add(key any) {
	q.arr = append(q.arr, key)
}

func (q *Queue) Take() (any, error) {
	switch {
	case len(q.arr) == 0:
		return nil, emptyQueueError{}
	default:
		elem := q.arr[0]
		q.arr = q.arr[1:len(q.arr)]
		return elem, nil
	}

}
