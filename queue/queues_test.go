package queue

import (
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	tqueue1 := NewLinkedQueue[int]()

	if !tqueue1.Empty() {
		t.Fatal("Empty (Case1): Not passed")
	}

	tqueue1.Enqueue(1)

	if tqueue1.Peak() != 1 {
		t.Fatal("Enqueue/Peak (Case1): Not passed, incorrect peak value")
	}

	if tqueue1.Empty() {
		t.Fatal("Empty (Case2): Not passed")
	}

	tqueue1.Enqueue(2)
	tqueue1.Enqueue(3)

	if tqueue1.Peak() != 1 {
		t.Fatal("Enqueue/Peak (Case2): Not passed, incorrect peak value")
	}

	val := tqueue1.Dequeue()

	if val != 1 {
		t.Error("Dequeue (Case1): Not passed, incorrect dequeue return")
	}

	if tqueue1.Peak() != 2 {
		t.Fatal("Dequeue (Case1.5): Not passed, next value incorrect")
	}

	tqueue1.Dequeue()
	tqueue1.Dequeue()

	if !tqueue1.Empty() {
		t.Fatal("Dequeue/Empty (Case1): not functioning properly")
	}
}