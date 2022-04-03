package main

import (
	"github.com/HenriMalahieude/GoStructures/bst"
)

func main() {
	bt1 := bst.NewBinarySearchTree(5, 
	func(a, b int) bool {
		return a > b
	}, 
	func (a, b int) bool {
		return a == b
	})

	bt1.Insert(3)
	bt1.Insert(4)
	bt1.Insert(1)
	bt1.Insert(2)
	bt1.Insert(0)

	bt1.Insert(7)
	bt1.Insert(6)
	bt1.Insert(8)

	bt1.Remove(3)

	bt1.InorderPrint()
	//fmt.Println()

	bt1.VisualizeDotty("bt")
}