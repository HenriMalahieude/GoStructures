package main

import (
	"fmt"

	"github.com/HenriMalahieude/GoStructures/queue"
)

func testQueue() {
	defer catch("Queue")

	var total *int = new(int)
	tqueue1 := queue.NewLinkedQueue[int]()


	doCheck(total, "Empty (Case1): Not passed", func() bool {
		return tqueue1.Empty()
	})

	tqueue1.Enqueue(1)

	doCheck(total, "Enqueue/Peak (Case1): Not passed, incorrect peak value", func() bool {
		return tqueue1.Peak() == 1
	})

	doCheck(total, "Empty (Case2): Not passed", func() bool {
		return !tqueue1.Empty()
	})

	tqueue1.Enqueue(2)
	tqueue1.Enqueue(3)

	doCheck(total, "Enqueue/Peak (Case2): Not passed, incorrect peak value", func() bool {
		return tqueue1.Peak() == 1
	})

	val := tqueue1.Dequeue()

	doCheck(total, "Dequeue (Case1): Not passed, incorrect dequeue return", func() bool {
		return val == 1
	})

	doCheck(total, "Dequeue (Case1.5): Not passed, next value incorrect", func() bool {
		return tqueue1.Peak() == 2
	})

	tqueue1.Dequeue()
	tqueue1.Dequeue()

	doCheck(total, "Dequeue/Empty (Case1): not functioning properly", func() bool {
		return tqueue1.Empty()
	})

	fmt.Print("Tests Passed on Queue: ", total, " / 7\n\n")
}