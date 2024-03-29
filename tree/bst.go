package tree

//NewBinarySearchTree Creates an empty Binary Search tree
func NewBinarySearchTree[T any](compareFunction, equals func(T, T) bool) *BinaryTree[T] {
	return &BinaryTree[T]{
		comparator: compareFunction,
		equalizer:  equals,
		root:       nil,
		nextID:     0,
	}
}

//Insert Inserts a new value into the tree, using the comparator function saved in the struct
func (b *BinaryTree[T]) Insert(value T) {
	if b.root == nil {
		b.root = new(binaryNode[T])
		b.root.id = 0
		b.root.value = value
		b.nextID = 1
		return
	}

	var insertAfter *binaryNode[T] = b.root
	for insertAfter != nil {
		if b.comparator(value, insertAfter.value) { //if true, go to the right
			if insertAfter.right == nil {
				break
			}
			insertAfter = insertAfter.right
		} else {
			if insertAfter.left == nil {
				break
			}
			insertAfter = insertAfter.left
		}
	}

	if b.comparator(value, insertAfter.value) {
		insertAfter.right = new(binaryNode[T])
		insertAfter.right.value = value
		insertAfter.right.id = b.nextID
		//insertAfter.right.name = name
		b.nextID++
	} else {
		insertAfter.left = new(binaryNode[T])
		insertAfter.left.value = value
		insertAfter.left.id = b.nextID
		//insertAfter.left.name = name
		b.nextID++
	}
}

//Remove the value from location, and replace with best replacement
func (b *BinaryTree[T]) Remove(value T) {
	//Locate the node we want to remove
	var remNode *binaryNode[T] = b.root
	var remPar *binaryNode[T] = remNode
	var remSide bool = false //true for right side

	for remNode != nil {
		if b.equalizer(value, remNode.value) {
			break
		}

		remPar = remNode
		if b.comparator(value, remNode.value) { //if true, go to the right
			remNode = remNode.right
			remSide = true
		} else {
			remNode = remNode.left
			remSide = false
		}
	}

	//fmt.Println(remNode.value, remPar.value)

	//Locate the replacement, and replace
	var repPar *binaryNode[T] = remPar   //replacement parent
	var repNode *binaryNode[T] = remNode //replacement node
	var repSide bool = remSide           //the side it's on

	if repNode.right == nil && repNode.left != nil { //right child doesn't exist
		//fmt.Println("Looking for replacement on left side")
		repPar = repNode
		repNode = repNode.left

		repSide = false

		for repNode != nil && repNode.right != nil {
			repPar = repNode
			repNode = repNode.right
			repSide = true
		}
	} else if repNode.left == nil && repNode.right != nil { //left child doesn't exist
		//fmt.Println("Looking for replacement on right side")
		repPar = repNode
		repNode = repNode.right

		repSide = true

		for repNode != nil && repNode.left != nil {
			repPar = repNode
			repNode = repNode.left
			repSide = false
		}
	}

	/*fmt.Println("Remove Node: " + fmt.Sprintf("%v", remNode.value))
	fmt.Println("Replacement Node: " + fmt.Sprintf("%v", repNode.value))
	fmt.Println(remPar.value, remPar.left.value, remPar.right.value)
	fmt.Println(remNode.value, remNode.left == nil, remNode.right == nil)
	fmt.Println(repPar.value, repPar.left.value, repPar.right.value)
	fmt.Println(repNode.value, repNode.left == nil, repNode.right == nil)*/

	//We have found the replacement and now will switch em out
	if repNode != nil {
		if (repNode.right != nil || repNode.left != nil) && !b.equalizer(repNode.value, value) {
			b.Remove(repNode.value)
		} else { //that means that replacement node has no children, and we can remove the "leaf" node from the tree
			if repSide {
				repPar.right = nil
			} else {
				repPar.left = nil
			}
		}

		repNode.right = remNode.right
		repNode.left = remNode.left

		if repNode == remNode { //if there is no good replacement (because of no children)
			if remSide {
				remPar.right = nil
			} else {
				remPar.left = nil
			}
			return
		}

		if remNode != b.root {
			if remSide {
				remPar.right = repNode
			} else {
				remPar.left = repNode
			}
		} else {
			b.root = repNode
		}
	} else {
		panic("Replacement Node was nil?")
	}
}

//Search returns the binary node containing the value given, nil otherwise
func (b *BinaryTree[T]) Search(value T) bool {
	curNode := b.root
	for curNode != nil {
		if b.equalizer(value, curNode.value) {
			return true
		}

		if b.comparator(value, curNode.value) { //if true, go to the right
			curNode = curNode.right
		} else {
			curNode = curNode.left
		}
	}

	return false
}
