package stack_test

import (
	"fmt"
	"testing"

	"github.com/Pawilonek/algo/stack"
)

func TestStackOrder(t *testing.T) {
	testingElements := []string{}
	for i := 0; i < 50; i++ {
		testingElements = append(testingElements, fmt.Sprintf("Elelent no. %02d", i))
	}

	s := stack.New()
	for _, elm := range testingElements {
		s.Push(elm)
	}

	for i := len(testingElements) - 1; i >= 0; i-- {
		expectedElement := testingElements[i]
		stackElement, _ := s.Pop()
		if expectedElement != stackElement {
			t.Errorf("Invalid order of stack elements. Expected %s, got: %s", expectedElement, stackElement)
			break
		}
	}
}

func TestOverPop(t *testing.T) {
	s := stack.New()

	_, err := s.Pop()
	if err == nil {
		t.Errorf("Expected error about empty stack")
	}
}

func TestIsEmpty(t *testing.T) {
	s := stack.New()

	if s.IsEmpty() != true {
		t.Errorf("New stack should be empty")

	}

	s.Push("test")
	if s.IsEmpty() != false {
		t.Errorf("Stack should not be empty with added new element")

	}

	s.Pop()
	if s.IsEmpty() != true {
		t.Errorf("Stack shoud be empty with all the elements pulled from queue")
	}
}
