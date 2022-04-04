package queue

//Queue encompasses all different implementations of a queue
type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Peak() T
	Empty() bool
}

type element[T any] struct {
	entry T
	next  *element[T]
}

//LinkedQueue is a linked list implementation of a queue
type LinkedQueue[T any] struct {
	front *element[T]
	back  *element[T]
}