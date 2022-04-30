package list

import "errors"

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

	d.head.prev = newNode
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

func (d *DoublyLinkedList[T]) Insert(pos uint, value T){
	if (d.insertionSpecialCase(value)){
		return
	}

	var curPos uint = 0;
	var insertAfter *DoubleNode[T] = d.head;
	for insertAfter != nil && insertAfter.next != nil && (curPos+1) < pos{
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

//Get returns the value of the 
func (d *DoublyLinkedList[T]) Get(pos uint) (T, error){
	var curPos uint = 0;
	var node *DoubleNode[T] = d.head;
	for node != nil && curPos < pos{
		node = node.next;
		curPos++;
	}

	if node == nil {
		defVal := new(T)

		return *defVal, errors.New("out of bounds access")
	}

	return node.entry, nil
}

func (d *DoublyLinkedList[T]) RemoveFirst(lambda func(T) bool) {
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
		d.head.prev = nil
	}else if (curNode == d.tail){
		d.tail = curNode.prev
		d.tail.next = nil
	}else{
		curNode.prev.next = curNode.next
	}
}

func (d *DoublyLinkedList[T]) Search(lambda func(T) bool) (uint, error) {
	var curPos uint = 0;
	var curNode *DoubleNode[T] = d.head;
	for curNode != nil && curNode.next != nil{
		if lambda(curNode.entry) {
			return curPos, nil
		}

		curPos++;
		curNode = curNode.next
	}

	return 0, errors.New("not found")
}

func (d *DoublyLinkedList[T]) Empty() bool {
	return d.head == nil
}