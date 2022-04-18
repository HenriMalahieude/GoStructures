package tree

import "testing"

func TestBinarySearchTree(t *testing.T) {
	tbst1 := NewBinarySearchTree(func(a, b int) bool {
		return a < b
	}, func(a, b int) bool {
		return a == b
	})

	tbst1.Insert(5)

	if !tbst1.Search(5) {
		t.Fatal("Insert/Search (Case1): Not Passed, root case")
	}

	tbst1.Insert(3)
	tbst1.Insert(7)

	if !tbst1.Search(3) && !tbst1.Search(7) {
		t.Fatal("Insert/Search (Case2): Not Passed, leaf case")
	}

	tbst1.Insert(2)
	tbst1.Insert(4)
	tbst1.Insert(6)
	tbst1.Insert(8)

	if !tbst1.Search(3) && !tbst1.Search(7) {
		t.Fatal("Insert/Search (Case3): Not Passed, internal case")
	}

	if !tbst1.Search(2) && !tbst1.Search(6) {
		t.Fatal("Insert/Search (Case4): Not Passed, leaf trip case")
	}

	tbst1.Remove(2)

	if tbst1.Search(2) {
		t.Fatal("Remove/Search (Case1): Not Passed, leaf case")
	}

	tbst1.Remove(7)

	if tbst1.Search(7) {
		t.Fatal("Remove/Search (Case2): Not Passed, leaf case")
	}

	tbst1.Remove(5)

	if tbst1.Search(5) {
		t.Fatal("Remove/Search (Case3): Not Passed, root case")
	}
}