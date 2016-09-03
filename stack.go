// stack package implements the stack with array
package stack

import "sync"

//
// stacks uses orded LIFO(Last In First Out) or FILO(First In Last Out).
//
type Stack struct {
	// v holds the stack values as an interface array
	// so, we can add any kind of item into the list.
	v []interface{}

	// count holds the stack lenght inside the Stack struct
	count int

	// mutex satisfy the synchronization for stack usage in goroutines
	mu sync.RWMutex
}

// NewStack inits the stack
func NewStack() *Stack {
	return &Stack{
		v: make([]interface{}, 0),
	}
}

// Push adds an item to the list(stack)
// Then increses count of the Stack
func (s *Stack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.v = append(s.v, v)
	s.count++
}

// Pop gets the last added item from the list
// Also descreases the count of stack
// if count isn't 0 already
func (s *Stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isEmpty() {
		return nil
	}

	v := s.v[s.count-1]
	s.v = s.v[:s.count-1]
	s.count--
	return v
}

// Peek gets the top value inside from stack
// it doesn't Pop the stack,
// Also, count of the stack won't be changed
func (s *Stack) Peek() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.isEmpty() {
		return nil
	}

	return s.v[len(s.v)-1]
}

// Count return the length of the stack
func (s *Stack) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.count

}

// isEmpty checks the stack if it has any item in it.
func (s *Stack) isEmpty() bool {
	if s.count == 0 {
		return true
	}
	return false
}
