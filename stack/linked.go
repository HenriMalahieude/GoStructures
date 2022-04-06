package stack

//NewLinkedStack returns an empty stack based on a linkedlist style
func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{
		nil,
	}
}

//Push adds a new element to the top of the stack
func (ls *LinkedStack[T]) Push(value T){
	if ls.top == nil {
		ls.top = new(element[T])
		ls.top.entry = value
		return
	}

	ls.top.next = new(element[T])
	ls.top.next.entry = value
	ls.top.next.prev = ls.top
	ls.top = ls.top.next
}

//Pop removes and returns the element at the top of the stack
func (ls *LinkedStack[T]) Pop() T {
	val := ls.top.entry
	ls.top = ls.top.prev
	if ls.top != nil {
		ls.top.next = nil
	}
	

	return val
}

//Peak returns the value stored in the top element
func (ls *LinkedStack[T]) Peak() T {
	return ls.top.entry
}

func (ls *LinkedStack[T]) Empty() bool {
	return ls.top == nil
}