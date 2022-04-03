package list

//NewSinglyLinkedList returns a new singly linked list
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T]{
	list := &SinglyLinkedList[T]{
		nil,
		nil,
	}

	return list
}

func (l *SinglyLinkedList[T]) insertionSpecialCase(value T) bool{
	if (l.head == nil){ //special case of empty
		l.head = &SingleNode[T]{
			value,
			nil,
		}
		l.tail = l.head
		
		return true
	}

	return false
}

//PushFront pushes the value at the front/head of the list
func (l *SinglyLinkedList[T]) PushFront(value T){
	if (l.insertionSpecialCase(value)){
		return
	}

	newNode := &SingleNode[T]{
		value,
		l.head, //I think this works? Right?
	}

	l.head = newNode
}

//PushBack pushes the value at the back/tail of the list
func (l *SinglyLinkedList[T]) PushBack(value T){
	if (l.insertionSpecialCase(value)){
		return
	}

	newNode := &SingleNode[T]{
		value,
		nil, //I think this works? Right?
	}

	l.tail.next = newNode
	l.tail = newNode
}

//Insert inserts the value into the list at position specified (or at end of list of too far)
func (l *SinglyLinkedList[T]) Insert(pos int, value T){
	if (l.insertionSpecialCase(value)){
		return
	}

	var curPos int = 0;
	var insertAfter *SingleNode[T] = l.head;
	for insertAfter != nil && insertAfter.next != nil && curPos < pos{
		insertAfter = insertAfter.next;
		curPos++;
	}

	if (insertAfter == l.head){
		l.PushFront(value)
	}else if (insertAfter == l.tail){
		l.PushBack(value)
	}else{
		newNode := &SingleNode[T]{
			value,
			insertAfter.next,
		}

		insertAfter.next = newNode
	}
}

//Search searches the list and returns the first entry which satisfies the function provided, nil otherwise
func (l *SinglyLinkedList[T]) Search(lambda func(T) bool) *SingleNode[T]{
	var curNode *SingleNode[T] = l.head;
	for curNode != nil && curNode.next != nil{
		if lambda(curNode.entry) {
			return curNode
		}

		curNode = curNode.next
	}

	return nil
}