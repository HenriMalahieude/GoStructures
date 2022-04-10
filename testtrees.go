package main

import (
	"fmt"

	"github.com/HenriMalahieude/GoStructures/tree"
)

func testBST() {
	defer catch("BST")

	var total *int = new(int)
	(*total) = 0

	tbst1 := tree.NewBinarySearchTree(func(a, b int) bool {
		return a < b
	}, func(a, b int) bool {
		return a == b
	})

	tbst1.Insert(5)

	doCheck(total, "Insert/Search (Case1): Not Passed, root case", func() bool {
		return tbst1.Search(5)
	})

	tbst1.Insert(3)
	tbst1.Insert(7)

	doCheck(total, "Insert/Search (Case2): Not Passed, leaf case", func() bool {
		return tbst1.Search(3) && tbst1.Search(7)
	})

	tbst1.Insert(2)
	tbst1.Insert(4)
	tbst1.Insert(6)
	tbst1.Insert(8)

	doCheck(total, "Insert/Search (Case3): Not Passed, internal case", func() bool {
		return tbst1.Search(3) && tbst1.Search(7)
	})

	doCheck(total, "Insert/Search (Case4): Not Passed, leaf trip case", func() bool {
		return tbst1.Search(2) && tbst1.Search(6)
	})

	tbst1.Remove(2)

	doCheck(total, "Remove/Search (Case1): Not Passed, leaf case", func() bool {
		return !tbst1.Search(2)
	})

	tbst1.Remove(7)

	doCheck(total, "Remove/Search (Case2): Not Passed, internal case", func() bool {
		return !tbst1.Search(7)
	})

	tbst1.Remove(5)

	doCheck(total, "Remove/Search (Case3): Not Passed, root case", func() bool {
		return !tbst1.Search(5)
	})

	//There may need more tests, but this is comprehensive enough atm
	fmt.Print("Test Passed on BST: ", *total, " / 7\n\n")
}