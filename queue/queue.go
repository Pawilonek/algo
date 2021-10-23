package queue

import (
	"fmt"
)

type Queue struct {
	queue []interface{}
}

func NewQueue() Queue {
	return Queue{
		queue: make([]interface{}, 0),
	}
}

func (q *Queue) Push(element interface{}) {
	q.queue = append(q.queue, element)
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return "", fmt.Errorf("The queue is empty")
	}

	element := q.queue[0]
	q.queue = q.queue[1:]

	return element, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
