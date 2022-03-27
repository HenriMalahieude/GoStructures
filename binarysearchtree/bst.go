package binarysearchtree

import (
	"fmt"
)

//BinaryNode Simple value as defined by user with two children
type BinaryNode[T any] struct {
	value T
	left, right  *BinaryNode[T]
}

//BinaryTree Simple Binary Tree, can define the comparator function (must be true for right value)
type BinaryTree[T any] struct {
	comparator func(T, T) bool
	root *BinaryNode[T]
}

//NewBinarySearchTree Creates a new Binary Search tree
func NewBinarySearchTree[T any](startValue T, compareFunction func(T, T) bool) *BinaryTree[T] {
	nRootNode := new(BinaryNode[T])
	nRootNode.value = startValue
	
	return &BinaryTree[T]{
		comparator: compareFunction,
		root: nRootNode,
	}
}

//Insert Inserts a new value into the tree, using the comparator function saved in the struct
func (b *BinaryTree[T]) Insert(value T){
	var insertAfter *BinaryNode[T] = b.root
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
		insertAfter.right = new(BinaryNode[T])
		insertAfter.right.value = value
	}else{
		insertAfter.left = new(BinaryNode[T])
		insertAfter.left.value = value
	}
}

//Remove the value from location, and replace with best replacement
func (b *BinaryTree[T]) Remove(value T, equal func(T, T) bool){
	//Locate the node we want to remove
	var removeNode *BinaryNode[T] = b.root
	for removeNode != nil {
		if equal(value, removeNode.value) { //if true, go to the right
			removeNode = removeNode.right
		}else{
			removeNode = removeNode.left
		}
	}

	//Locate the replacement
	var replacement *BinaryNode[T] = removeNode
	if replacement.right == nil {
		
	}
}

//InorderPrint Outputs the tree in order of seeing each node
func (b BinaryTree[T]) InorderPrint(){
	recurseInOrder(b.root)
}

func recurseInOrder[T any](node *BinaryNode[T]){
	if node == nil {return}

	recurseInOrder(node.left)
	fmt.Printf("%v ", node.value)
	recurseInOrder(node.right)
}

//PreorderPrint Outputs the tree with the root node coming first
func (b BinaryTree[T]) PreorderPrint(){

}

//PostorderPrint Outputs the tree with the root node coming last
func (b BinaryTree[T]) PostorderPrint(){

}