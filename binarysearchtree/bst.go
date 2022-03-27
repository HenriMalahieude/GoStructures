package binarysearchtree

import "fmt"

//BinaryNode Simple value as defined by user with two children
type BinaryNode struct {
	value interface{}
	left  *BinaryNode
	right *BinaryNode
}

//BinaryTree Simple Binary Tree, can define the comparator function (must be true for right value)
type BinaryTree struct {
	comparator func(interface{}, interface{}) bool
	root *BinaryNode
}

//NewBinarySearchTree Creates a new Binary Search tree
func NewBinarySearchTree(startValue interface{}, compareFunction func(interface{}, interface{}) bool) *BinaryTree {
	nRootNode := new(BinaryNode)
	nRootNode.value = startValue
	
	return &BinaryTree{
		comparator: compareFunction,
		root: nRootNode,
	}
}

//Insert Inserts a new value into the tree, using the comparator function saved in the struct
func (b *BinaryTree) Insert(value interface{}){
	valType := fmt.Sprintf("%T", value)
	rooType := fmt.Sprintf("%T", b.root.value)
	if valType != rooType {
		panic("Attempted to insert a " + valType + " into a " + rooType + " binary search tree.") //you handle it yourself baby, not my problem
	}

	var insertAfter *BinaryNode = b.root
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
		insertAfter.right = new(BinaryNode)
		insertAfter.right.value = value
	}else{
		insertAfter.left = new(BinaryNode)
		insertAfter.left.value = value
	}
}

//Remove removes the 
func (b *BinaryTree) Remove(value interface{}){

}

//InorderPrint Outputs the tree in order of seeing each node
func (b BinaryTree) InorderPrint(){
	recurseInOrder(b.root)
}

func recurseInOrder(node *BinaryNode){
	if node == nil {return}

	recurseInOrder(node.left)
	fmt.Printf("%v ", node.value)
	recurseInOrder(node.right)
}

//PreorderPrint Outputs the tree with the root node coming first
func (b BinaryTree) PreorderPrint(){

}

//PostorderPrint Outputs the tree with the root node coming last
func (b BinaryTree) PostorderPrint(){

}