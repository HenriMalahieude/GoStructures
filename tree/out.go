package tree

import "fmt"

//InorderPrint Outputs the tree in order of seeing each node
func (b BinaryTree[T]) InorderPrint() {
	recurseInOrder(b.root)
}

func recurseInOrder[T any](node *binaryNode[T]) {
	if node == nil {
		return
	}

	recurseInOrder(node.left)
	fmt.Printf("%v ", node.value)
	recurseInOrder(node.right)
}

//PreorderPrint Outputs the tree with the root node coming first
func (b BinaryTree[T]) PreorderPrint() {
	recursePreOrder(b.root)
}

func recursePreOrder[T any](node *binaryNode[T]) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.value)
	recursePreOrder(node.left)
	recursePreOrder(node.right)
}

//PostorderPrint Outputs the tree with the root node coming last
func (b BinaryTree[T]) PostorderPrint() {
	recursePostOrder(b.root)
}

func recursePostOrder[T any](node *binaryNode[T]) {
	if node == nil {
		return
	}

	recursePostOrder(node.left)
	recursePostOrder(node.right)
	fmt.Printf("%v ", node.value)
}
