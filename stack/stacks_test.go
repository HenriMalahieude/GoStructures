package stack

import "testing"

//TestStack tests stack
func TestLinkedStack(t *testing.T) {
	tstack1 := NewLinkedStack[int]()

	if !tstack1.Empty() {
		t.Fatal("Empty (Case1): Not Passed")
	}

	tstack1.Push(1)

	if tstack1.Peak() != 1 {
		t.Fatal("Push/Peak (Case1): Not Passed, did not push/peak correctly")
	}

	if tstack1.Empty() {
		t.Fatal("Empty (Case2): Not passed")
	}

	tstack1.Push(2)
	tstack1.Push(3)

	if tstack1.Peak() != 3 {
		t.Error("Push/Peak (Case2): Not Passed, did not push/peak correctly")
	}

	rem := tstack1.Pop()

	if tstack1.Peak() != 2 {
		t.Fatal("Pop/Peak (Case1): Not Passed, did not pop/peak correctly")
	}

	if rem != 3 {
		t.Error("Pop (Case1): Not Passed, incorrect return")
	}

	tstack1.Pop()
	tstack1.Pop()

	if !tstack1.Empty() {
		t.Fatal("Empty/Pop (Case1): Not Passed")
	}
}