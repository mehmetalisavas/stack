# Stack

Supported with concurrency written in Go.
You can use stack implementation supported with goroutines easily.

# Stack Description

Stack is a linear data structure which follows a particular order in which the operations are performed.
The order may be LIFO(Last In First Out) or FILO(First In Last Out).

Three basic operations are performed in the stack:
- Push: Adds an item in the stack. And Increases count of Stack
- Pop: Removes an item from the stack.If the stack is empty, then it return nil value. And Decreases count of Stack
- Peek: Get the topmost item. Count if Stack won't be changed when Peek is called

# Example

```go
	stack := NewStack() //creates a new stack

	stack.Push("data") // pushes into the stack 'data' value (stack count size : 1)
	stack.Push(1) // pushes into the stack '1' value (stack count size : 2)
	stack.Push(true) // pushes into the stack 'true' value (stack count size : 3)

	data := stack.Pop() // returns the last value 'true' , in stack count size : 2
	data := stack.Peek() // returns the top value '1' (stack count size : 2)- count won't be changed

```
