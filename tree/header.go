package tree

//Tree is any BinarySearch, AVL, or Balance Tree
type Tree[T any] interface {
	Insert(value T)
	Remove(value T)
	Search(value T) bool
}

//BinaryNode Simple value as defined by user with two children
type BinaryNode[T any] struct {
	value       T
	left, right *BinaryNode[T]
	id          int
	//name string
}

//BinaryTree Simple Binary Tree, must define the comparator function (must be true for larger value), and equalizer function
type BinaryTree[T any] struct {
	comparator func(T, T) bool
	equalizer  func(T, T) bool
	root       *BinaryNode[T]
	nextID     int
}