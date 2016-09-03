package stack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack.count != 0 {
		t.Fatal("stack count should equal zero")
	}
	if len(stack.v) != 0 {
		t.Fatal("stack length should equal zero")
	}
}

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stack.Push("test")
	if stack.count == 0 {
		t.Fatal("stack count should not equal zero")
	}
	if stack.count != 1 {
		t.Fatal("stack count should equal one")
	}
	if len(stack.v) == 0 {
		t.Fatal("stack length should not equal zero")
	}

	if len(stack.v) != 1 {
		t.Fatal("stack length should equal one")
	}
}

func TestStackPop(t *testing.T) {
	stack := NewStack()
	data := stack.Pop()
	if data != nil {
		t.Fatal("stack data should be nil")
	}

	stack.Push("data")
	data = stack.Pop()
	if data == nil {
		t.Fatal("stack data should not be nil")
	}
	if stack.count != 0 {
		t.Fatal("stack length should equal zero")
	}

}

func TestStackPeek(t *testing.T) {
	stack := NewStack()
	stack.Push("data")
	data := stack.Peek()
	if data == nil {
		t.Fatal("stack data should not be nil")
	}
	if stack.count != 1 {
		t.Fatal("stack length should equal one")
	}

	_ = stack.Pop()
	data = stack.Peek()
	if data != nil {
		t.Fatal("stack data should be nil")
	}
	if stack.count != 0 {
		t.Fatal("stack length should equal zero")
	}
}

func TestCount(t *testing.T) {
	stack := NewStack()
	if stack.Count() != 0 {
		t.Fatal("stack count should be zero")
	}

	stack.Push("data")
	stack.Push("data")
	if stack.Count() != 2 {
		t.Fatal("stack count should be 2")
	}
	stack.Pop()
	if stack.Count() != 1 {
		t.Fatal("stack count should be 1")
	}

}

func TestPushPop(t *testing.T) {
	stack := NewStack()
	stack.Push("data")
	stack.Push("data")
	stack.Push("data")
	_ = stack.Pop()

	if stack.count != 2 {
		t.Fatal("stack count should be 2")
	}
	if stack.v[0] != "data" {
		t.Fatal("First item of stack should be 'data'")
	}

}
