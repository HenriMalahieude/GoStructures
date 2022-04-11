package list

import "errors"

//NewCircularList returns a new empty circular list
func NewCircularList[T any]() *CircularList[T] {
	return &CircularList[T]{
		nil,
	}
}

//PushFront inserts an element as the new head, therefore at the "front" of the list
func (c *CircularList[T]) PushFront(value T){
	c.PushBack(value)
	
	c.head = c.head.prev
}

//PushBack inserts an element behind the head
func (c *CircularList[T]) PushBack(value T){
	if c.head == nil {
		c.head = new(DoubleNode[T])
		c.head.entry = value
		c.head.next = c.head
		return
	}

	newNode := new(DoubleNode[T])
	newNode.entry = value
	newNode.prev = c.head.prev
	newNode.next = c.head

	c.head.prev.next = newNode
}

func (c *CircularList[T]) Insert(pos uint, value T){
	if c.head == nil {
		c.head = new(DoubleNode[T])
		c.head.entry = value
		c.head.prev = c.head
		c.head.next = c.head
		return
	}

	var curPos uint = 0;
	var insertAfter *DoubleNode[T] = c.head;
	for curPos < pos{ //not worried about nil things
		insertAfter = insertAfter.next;
		curPos++;
	}

	if (insertAfter == c.head){
		c.PushFront(value)
	}else if (insertAfter.next == c.head){
		c.PushBack(value)
	}else{
		newNode := &DoubleNode[T]{
			value,
			insertAfter,
			insertAfter.next,
		}

		insertAfter.next = newNode
	}
}

func (c *CircularList[T]) Get(pos uint) (T, error){
	if c.head == nil {
		defVal := new(T)
		return *defVal, errors.New("empty circle")
	}

	var curPos uint = 0;
	var node *DoubleNode[T] = c.head;
	for curPos < pos{
		node = node.next;
		curPos++;
	}

	return node.entry, nil
}

func (c *CircularList[T]) Remove(lambda func(T) bool){
	if c.head == nil{return}

	remHead := true
	curNode := c.head
	if !lambda(curNode.entry){
		remHead = false
		curNode = curNode.next

		for curNode != c.head {
			if lambda(curNode.entry) {
				break;
			}
			curNode = curNode.next
		}
	}

	if (curNode == c.head && !remHead){ //the
		return
	}

	curNode.prev.next = curNode.next	
}

func (c *CircularList[T]) Search(lambda func(T) bool) (uint, error){
	isHead := true
	var curPos uint = 0
	curNode := c.head
	if !lambda(curNode.entry){
		isHead = false
		curNode = curNode.next
		curPos++;

		for curNode != c.head {
			if lambda(curNode.entry) {
				break;
			}
			curPos++;
			curNode = curNode.next
		}
	}

	if curNode == c.head && !isHead {
		return 0, errors.New("not found")
	}

	return curPos, nil
}

func (c *CircularList[T]) Empty() bool {
	return c.head == nil
}