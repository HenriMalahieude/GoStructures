package list

//LinkedList represents the different types of lists we can have
type LinkedList[T any] interface {
	PushFront(T)
	PushBack(T)
	Insert(pos int, value T)
	Remove(func(T) bool)
	Search(func(T) bool) any
}

//SingleNode is a node for a singly linked list
type SingleNode[T any] struct {
	entry T
	next *SingleNode[T]
}

//SinglyLinkedList is a list where each node only connects to the next
type SinglyLinkedList[T any] struct {
	head *SingleNode[T]
	tail *SingleNode[T]
}

//DoubleNode is a node for a doubly linked list
type DoubleNode[T any] struct {
	entry T
	back *DoubleNode[T]
	next *DoubleNode[T]
}

//DoublyLinkedList is a list where each node connects to its next and previous node
type DoublyLinkedList[T any] struct{
	head *DoubleNode[T]
	tail *DoubleNode[T]
}