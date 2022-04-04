package stack

//Stack is any stack type
type Stack[T any] interface {
	Push(T);
	Pop() T;
	Peak() T;
	Empty() bool;
}

type element[T any] struct {
	entry T
	prev *element[T]
	next *element[T]
}

//LinkedStack is linked list style stack
type LinkedStack[T any] struct {
	top *element[T]
}

//TODO: Array-based Stack