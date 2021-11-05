package stack

import (
	"fmt"
)

// Stack is an implementation of a Stack (First In Last Out queue)
type Stack struct {
	stack []interface{}
	top   int
}

// New creates new instance of Stack
func New() Stack {
	return Stack{
		stack: make([]interface{}, 0),
		top:   -1,
	}
}

// Push adds new element on stack
func (s *Stack) Push(element interface{}) {
	s.stack = append(s.stack, element)
	s.top++
}

// Pop returns latest element from the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("The stack is empty")
	}

	element := s.stack[s.top]
	s.stack = s.stack[:s.top]
	s.top--

	return element, nil
}

// IsEmpty reruns information if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top < 0
}
