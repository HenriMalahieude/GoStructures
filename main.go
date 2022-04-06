package main

import (
	"fmt"

	"github.com/HenriMalahieude/GoStructures/stack"
)

func testBST() {

}

func testStack() {
	var total int = 0;

	tstack1 := stack.NewLinkedStack[int]()

	if !tstack1.Empty(){
		fmt.Println("Empty (Case1): Not Passed")
	}else{
		total++
	}

	tstack1.Push(1)

	if tstack1.Peak() != 1 {
		fmt.Println("Push/Peak (Case1): Not Passed, did not push/peak correctly")
	}else{
		total++
	}

	tstack1.Push(2)
	tstack1.Push(3)

	if tstack1.Peak() != 3 {
		fmt.Println("Push/Peak (Case2): Not Passed, did not push/peak correctly")
	}else{
		total++
	}

	rem := tstack1.Pop()

	if tstack1.Peak() != 2 {
		fmt.Println("Pop/Peak (Case1): Not Passed, did not pop/peak correctly")
	}else{
		total++
	}

	if rem != 3 {
		fmt.Println("Pop (Case1): Not Passed, incorrect return")
	}else{
		total++
	}

	tstack1.Pop()
	tstack1.Pop()

	if !tstack1.Empty(){
		fmt.Println("Empty (Case2): Not Passed")
	}else{
		total++
	}

	fmt.Println("Test Passed on Stack: ", total, "/ 6")
}

func testQueue() {

}

func testList() {

}

func main() {
	testStack()
}