package list

//NewDoublyLinkedList returns an empty DoublyLinkedList
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		nil,
		nil,
	}
}

func (d *DoublyLinkedList[T]) insertionSpecialCase(value T) bool{
	if d.head == nil{
		d.head = &DoubleNode[T]{
			value,
			nil, 
			nil,
		}
		d.tail = d.head

		return true
	}
	return false
}

func (d *DoublyLinkedList[T]) PushFront(value T){
	if d.insertionSpecialCase(value){
		return
	}

	newNode := &DoubleNode[T]{
		value,
		nil,
		d.head,
	}

	d.head.back = newNode
	d.head = newNode
}

func (d *DoublyLinkedList[T]) PushBack(value T){
	if d.insertionSpecialCase(value){
		return
	}

	newNode := &DoubleNode[T]{
		value,
		d.tail,
		nil,
	}

	d.tail.next = newNode
	d.tail = newNode
}

func (d *DoublyLinkedList[T]) Insert(pos int, value T){
	if (d.insertionSpecialCase(value)){
		return
	}

	var curPos int = 0;
	var insertAfter *DoubleNode[T] = d.head;
	for insertAfter != nil && insertAfter.next != nil && curPos < pos{
		insertAfter = insertAfter.next;
		curPos++;
	}

	if (insertAfter == d.head){
		d.PushFront(value)
	}else if (insertAfter == d.tail){
		d.PushBack(value)
	}else{
		newNode := &DoubleNode[T]{
			value,
			insertAfter,
			insertAfter.next,
		}

		insertAfter.next = newNode
	}
}

func (d *DoublyLinkedList[T]) Remove(lambda func(T) bool) {
	if d.head == nil{return}

	curNode := d.head
	for curNode != nil && curNode.next != nil {
		if lambda(curNode.entry) {
			break;
		}
		curNode = curNode.next
	}

	if (curNode == nil){ //the 
		return
	}else if (curNode == d.head){
		d.head = curNode.next
		d.head.back = nil
	}else if (curNode == d.tail){
		d.tail = curNode.back
		d.tail.next = nil
	}else{
		curNode.back.next = curNode.next
	}
}

func (d *DoublyLinkedList[T]) Search(lambda func(T) bool) any {
	var curNode *DoubleNode[T] = d.head;
	for curNode != nil && curNode.next != nil{
		if lambda(curNode.entry) {
			return curNode
		}

		curNode = curNode.next
	}

	return nil
}