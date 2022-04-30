package list

//LinkedList represents the different types of lists we can have
type LinkedList[T any] interface {
	PushFront(T)
	PushBack(T)
	Insert(pos uint, value T)
	Get(uint) (T, error)
	RemoveFirst(func(T) bool)
	Search(func(T) bool) (uint, error)
	Empty() bool
}

//SingleNode is a node for a singly linked list
type SingleNode[T any] struct {
	entry T
	next  *SingleNode[T]
}

//SinglyLinkedList is a list where each node only connects to the next
type SinglyLinkedList[T any] struct {
	head *SingleNode[T]
	tail *SingleNode[T]
}

//DoubleNode is a node for a doubly linked list
type DoubleNode[T any] struct {
	entry T
	prev  *DoubleNode[T]
	next  *DoubleNode[T]
}

//DoublyLinkedList is a list where each node connects to its next and previous node
type DoublyLinkedList[T any] struct {
	head *DoubleNode[T]
	tail *DoubleNode[T]
}

//CircularList where the tail points to the head
type CircularList[T any] struct {
	head *DoubleNode[T]
}