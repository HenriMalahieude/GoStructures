package main

import (
	"fmt"

	"github.com/HenriMalahieude/bst"
)

func main() {
	bt1 := bst.NewBinarySearchTree(5, func(a, b int) bool {
		return a > b
	})

	bt1.Insert(3)
	bt1.Insert(4)
	bt1.Insert(2)

	bt1.Insert(7)
	bt1.Insert(6)
	bt1.Insert(8)

	bt1.InorderPrint()
	fmt.Println()
	bt1.PreorderPrint()
	fmt.Println()
	bt1.PostorderPrint()
	fmt.Println()
}