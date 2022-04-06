package main

import (
	"fmt"

	"github.com/HenriMalahieude/GoStructures/queue"
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

	if tstack1.Empty() {
		fmt.Println("Empty (Case2): Not passed")
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
		fmt.Println("Empty/Pop (Case1): Not Passed")
	}else{
		total++
	}

	fmt.Println("Test Passed on Stack: ", total, "/ 7")
}

func testQueue() {
	var total int = 0;
	tqueue1 := queue.NewLinkedQueue[int]()

	if !tqueue1.Empty() {
		fmt.Println("Empty (Case1): Not passed")
	}else{
		total++
	}

	tqueue1.Enqueue(1)

	if tqueue1.Peak() != 1 {
		fmt.Println("Enqueue/Peak (Case1): Not passed, incorrect peak value")
	}else{
		total++
	}

	if tqueue1.Empty() {
		fmt.Println("Empty (Case2): Not passed")
	}else{
		total++
	}

	tqueue1.Enqueue(2)
	tqueue1.Enqueue(3)

	if tqueue1.Peak() != 1 {
		fmt.Println("Enqueue/Peak (Case2): Not passed, incorrect peak value")
	}else{
		total++
	}

	val := tqueue1.Dequeue()

	if val != 1 {
		fmt.Println("Dequeue (Case1): Not passed, incorrect dequeue return")
	}else{
		total++
	}

	if tqueue1.Peak() != 2 {
		fmt.Println("Dequeue (Case1.5): Not passed, next value incorrect")
	}else{
		total++
	}

	tqueue1.Dequeue()
	tqueue1.Dequeue()

	if !tqueue1.Empty() {
		fmt.Println("Dequeue/Empty (Case1): not functioning properly")
	}else{
		total++
	}

	fmt.Println("Tests Passed on Queue:", total, "/ 7")
}

func testList() {

}

func main() {
	testStack()
	testQueue()
}