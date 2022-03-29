package bst

//binaryNode Simple value as defined by user with two children
type binaryNode[T any] struct {
	value T
	left, right  *binaryNode[T]
	ID int
}

//BinaryTree Simple Binary Tree, can define the comparator function (must be true for right value)
type BinaryTree[T comparable] struct {
	comparator func(T, T) bool
	root *binaryNode[T]
	nextID int
}

//NewBinarySearchTree Creates a new Binary Search tree
func NewBinarySearchTree[T comparable](startValue T, compareFunction func(T, T) bool) *BinaryTree[T] {
	nRootNode := new(binaryNode[T])
	nRootNode.value = startValue
	nRootNode.ID = 0
	
	return &BinaryTree[T]{
		comparator: compareFunction,
		root: nRootNode,
		nextID: 1,
	}
}

//Insert Inserts a new value into the tree, using the comparator function saved in the struct
func (b *BinaryTree[T]) Insert(value T){
	var insertAfter *binaryNode[T] = b.root
	for insertAfter != nil {
		if b.comparator(value, insertAfter.value) { //if true, go to the right
			if insertAfter.right == nil {
				break
			}
			insertAfter = insertAfter.right
		}else{
			if insertAfter.left == nil {
				break
			}
			insertAfter = insertAfter.left
		}
	}

	if b.comparator(value, insertAfter.value) {
		insertAfter.right = new(binaryNode[T])
		insertAfter.right.value = value
		insertAfter.right.ID = b.nextID
		b.nextID++
	}else{
		insertAfter.left = new(binaryNode[T])
		insertAfter.left.value = value
		insertAfter.left.ID = b.nextID
		b.nextID++
	}
}

//Remove the value from location, and replace with best replacement
func (b *BinaryTree[T]) Remove(value T,){
	//Locate the node we want to remove
	var removeNode *binaryNode[T] = b.root
	for removeNode != nil {
		if value == removeNode.value { //if true, go to the right
			removeNode = removeNode.right
		}else{
			removeNode = removeNode.left
		}
	}

	//Locate the replacement
	var par *binaryNode[T] = removeNode
	var replacement *binaryNode[T] = removeNode
	if replacement.right == nil {
		replacement = replacement.left
		for replacement != nil && replacement.right != nil {
			par = replacement
			replacement = replacement.right
		}
		removeNode.value = replacement.value
		par.right = nil

		return
	}

	replacement = replacement.right
	for replacement != nil && replacement.left != nil {
		par = replacement
		replacement = replacement.left
	}
	removeNode.value = replacement.value
	par.left = nil
}