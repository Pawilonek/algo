package queue_test

import (
	"fmt"
	"testing"

	"github.com/Pawilonek/algo/queue"
)

func TestQueueOrder(t *testing.T) {
	testingElements := []string{}
	for i := 0; i < 50; i++ {
		testingElements = append(testingElements, fmt.Sprintf("Elelent no. %02d", i))
	}

	q := queue.NewQueue()

	for _, elm := range testingElements {
		q.Push(elm)
	}

	for _, expectedElement := range testingElements {
		queueElement, _ := q.Pop()
		if expectedElement != queueElement {
			t.Errorf("Invalid order of queue elements. Expected %s, got: %s", expectedElement, queueElement)
			break
		}
	}
}

func TestOverPop(t *testing.T) {
	q := queue.NewQueue()

	_, err := q.Pop()
	if err == nil {
		t.Errorf("Expected error about empty queue")
	}
}

func TestIsEmpty(t *testing.T) {
	q := queue.NewQueue()

	if q.IsEmpty() != true {
		t.Errorf("New queues should be empty")

	}

	q.Push("test")
	if q.IsEmpty() != false {
		t.Errorf("Queue should not be empty with added new element")

	}

	q.Pop()
	if q.IsEmpty() != true {
		t.Errorf("Queue shoud be empty with all the elements pulled from queue")
	}
}
