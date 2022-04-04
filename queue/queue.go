package queue

//NewLinkedQueue returns an empty queue
func NewLinkedQueue[T any]() *LinkedQueue[T]{
	return &LinkedQueue[T]{
		nil,
		nil,
	}
}

//Enqueue at back of list
func (q *LinkedQueue[T]) Enqueue(value T) {
	if q.front == nil {
		q.front = new(element[T])
		q.front.entry = value
		q.back = q.front
		return;
	}

	q.back.next = new(element[T])
	q.back.next.entry = value
	q.back = q.back.next
}

//Dequeue removes and returns the first element in the queue
func (q *LinkedQueue[T]) Dequeue() T {
	val := q.front.entry
	q.front = q.front.next

	return val
}

//Peak returns the first element in the queue without removing it
func (q *LinkedQueue[T]) Peak() T {
	return q.front.entry
}

//Empty returns true if empty. Please check this before using dequeue or peak
func (q *LinkedQueue[T]) Empty() bool {
	return q.front == nil
}