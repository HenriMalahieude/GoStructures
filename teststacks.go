package main

import (
	"fmt"

	"github.com/HenriMalahieude/GoStructures/stack"
)

func testStack() {
	defer catch("Stack")

	var total int = 0

	tstack1 := stack.NewLinkedStack[int]()

	if !tstack1.Empty() {
		fmt.Println("Empty (Case1): Not Passed")
	} else {
		total++
	}

	tstack1.Push(1)

	if tstack1.Peak() != 1 {
		fmt.Println("Push/Peak (Case1): Not Passed, did not push/peak correctly")
	} else {
		total++
	}

	if tstack1.Empty() {
		fmt.Println("Empty (Case2): Not passed")
	} else {
		total++
	}

	tstack1.Push(2)
	tstack1.Push(3)

	if tstack1.Peak() != 3 {
		fmt.Println("Push/Peak (Case2): Not Passed, did not push/peak correctly")
	} else {
		total++
	}

	rem := tstack1.Pop()

	if tstack1.Peak() != 2 {
		fmt.Println("Pop/Peak (Case1): Not Passed, did not pop/peak correctly")
	} else {
		total++
	}

	if rem != 3 {
		fmt.Println("Pop (Case1): Not Passed, incorrect return")
	} else {
		total++
	}

	tstack1.Pop()
	tstack1.Pop()

	if !tstack1.Empty() {
		fmt.Println("Empty/Pop (Case1): Not Passed")
	} else {
		total++
	}

	fmt.Print("Test Passed on Stack: ", total, " / 7\n\n")
}